package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/bond"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/router"
)

// CSRFTokenMiddleware publishes the CSRF token under the "csrfToken"
// root-template variable so app.go.html can stamp it into the
// <meta name="csrf-token"> tag. Must run AFTER the framework CSRF
// middleware, which attaches the request-scoped token cache.
//
// The published value is not the token itself but a csrfMetaToken that
// reads csrf.TokenForRequest when the root template prints it, i.e. at
// render time. The token is resolved late on purpose: this middleware
// runs in the web stack, before route-level middleware such as
// auth.AuthMiddleware, and a remember-me revival there regenerates the
// session and rotates its CSRF token. A value read here would still name
// the replaced session, so the meta tag would disagree with the
// XSRF-TOKEN cookie and the csrf_token page prop and a submit using it
// would answer 419. Read at render time, the meta tag is byte-identical
// to every other reader on the response (the request-scoped cache follows
// the session id), on a revival, a sign-in, a sign-out and an anonymous
// visit alike.
//
// Publishes via bond.WithTemplateData, the only API bond's renderer
// reads when building the root-template Execute map. gonertia's own
// SetTemplateDatum lives in a different context key and never reaches
// the bond template.
func CSRFTokenMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		ctx := bond.WithTemplateData(c.Request.Context(), "csrfToken", csrfMetaToken{r: c.Request})
		c.Request = c.Request.WithContext(ctx)
		return next(c)
	}
}

// csrfMetaToken is the "csrfToken" root-template value. html/template
// prints a fmt.Stringer through its String method, so {{ .csrfToken }}
// resolves the token when the template executes. r is the request as the
// CSRF middleware left it: it carries the request-scoped token cache and
// the session holder a later revival or sign-in updates in place, so a
// read at render time sees the session the response is served under.
// No session or no token state renders an empty content attribute, as
// before.
type csrfMetaToken struct {
	r *http.Request
}

// String returns the request's CSRF token in its masked emission form,
// or "" when the request has none.
func (t csrfMetaToken) String() string {
	token, err := csrf.TokenForRequest(t.r)
	if err != nil {
		return ""
	}
	return token
}
