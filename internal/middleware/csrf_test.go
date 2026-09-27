package middleware

import (
	"html"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/velocitykode/velocity/bond"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/router"
)

// metaTokenPattern captures the content attribute of the root template's
// <meta name="csrf-token"> tag.
var metaTokenPattern = regexp.MustCompile(`<meta name="csrf-token" content="([^"]*)">`)

// propTokenPattern captures the csrf_token page prop from the page object
// bond embeds in the body.
var propTokenPattern = regexp.MustCompile(`"csrf_token":"([^"]*)"`)

// sessionSwitch is a CSRF SessionIDResolver whose answer a handler can
// change mid-request, the way a remember-me revival, a sign-in or a
// sign-out replaces the session the response is served under after the
// web middleware stack already ran.
type sessionSwitch struct {
	mu sync.Mutex
	id string
}

func (s *sessionSwitch) set(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.id = id
}

func (s *sessionSwitch) resolve(*http.Request) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.id == "" {
		return "", csrf.ErrNoSession
	}
	return s.id, nil
}

// rootTemplateBond builds a bond renderer over the application's real
// root template, with the asset helpers stubbed out, and a share-props
// function that publishes csrf_token the way AppModule does.
func rootTemplateBond(t *testing.T) *bond.Bond {
	t.Helper()
	root, err := os.ReadFile("../../resources/views/app.go.html")
	if err != nil {
		t.Fatalf("read root template: %v", err)
	}
	b, err := bond.New(bond.Config{
		RootTemplate: string(root),
		Funcs: template.FuncMap{
			"vite": func(string) template.HTML { return "" },
		},
	})
	if err != nil {
		t.Fatalf("bond.New: %v", err)
	}
	b.SetSharePropsFunc(func(r *http.Request) (bond.Props, error) {
		props := bond.Props{}
		if token, err := csrf.TokenForRequest(r); err == nil && token != "" {
			props["csrf_token"] = token
		}
		return props, nil
	})
	return b
}

// TestCSRFTokenMiddleware_MetaFollowsSessionAtRender runs a full-page GET
// through the web stack's CSRF middleware and CSRFTokenMiddleware, then
// switches the session inside the handler (as route-level auth middleware
// does on a remember-me revival) before rendering the root template. The
// rendered <meta name="csrf-token"> must equal the csrf_token prop and
// csrf.TokenForRequest at render time, never the token of the session the
// request started under.
func TestCSRFTokenMiddleware_MetaFollowsSessionAtRender(t *testing.T) {
	cases := []struct {
		name     string
		before   string // session id when the middleware stack runs
		after    string // session id when the handler renders
		wantMeta bool
	}{
		{name: "remember-me revival replaces the session", before: "session-before-revival", after: "session-after-revival", wantMeta: true},
		{name: "signed in, same session", before: "session-signed-in", after: "session-signed-in", wantMeta: true},
		{name: "session created during the request", before: "", after: "session-new", wantMeta: true},
		{name: "sign-out drops the session", before: "session-signed-in", after: "", wantMeta: false},
		{name: "anonymous", before: "", after: "", wantMeta: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sessions := &sessionSwitch{}
			sessions.set(tc.before)
			cfg := csrf.DefaultConfig()
			cfg.SessionIDResolver = sessions.resolve
			protector, err := csrf.NewE(cfg)
			if err != nil {
				t.Fatalf("csrf.NewE: %v", err)
			}
			b := rootTemplateBond(t)

			var earlyToken, renderToken string
			handler := func(c *router.Context) error {
				earlyToken, _ = csrf.TokenForRequest(c.Request)
				sessions.set(tc.after)
				if err := b.Render(c.Response, c.Request, "Dashboard", bond.Props{}); err != nil {
					return err
				}
				renderToken, _ = csrf.TokenForRequest(c.Request)
				return nil
			}
			chain := protector.RouterMiddleware()(CSRFTokenMiddleware(handler))

			c, rec := router.NewTestContext(http.MethodGet, "/dashboard")
			if err := chain(c); err != nil {
				t.Fatalf("chain: %v", err)
			}

			m := metaTokenPattern.FindStringSubmatch(rec.Body.String())
			if m == nil {
				t.Fatalf("no csrf-token meta tag in body:\n%s", rec.Body.String())
			}
			meta := html.UnescapeString(m[1])
			prop := ""
			if p := propTokenPattern.FindStringSubmatch(rec.Body.String()); p != nil {
				prop = p[1]
			}

			if !tc.wantMeta {
				if meta != "" || prop != "" {
					t.Fatalf("meta csrf-token = %q, csrf_token prop = %q, want both empty (no session at render)", meta, prop)
				}
				return
			}
			if renderToken == "" {
				t.Fatal("csrf.TokenForRequest at render returned no token")
			}
			if meta != renderToken {
				t.Fatalf("meta csrf-token = %q, want the render-time token %q (token before the session switch: %q)", meta, renderToken, earlyToken)
			}
			if prop != meta {
				t.Fatalf("csrf_token prop = %q, want it equal to the meta csrf-token %q", prop, meta)
			}
		})
	}
}
