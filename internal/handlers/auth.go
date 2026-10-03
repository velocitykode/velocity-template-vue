package handlers

import (
	"{{MODULE_NAME}}/internal/models"

	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/validation"
	"github.com/velocitykode/velocity/validation/vform"
	"github.com/velocitykode/velocity/view"
)

// LoginRequest is the form-request schema for POST /login. vform.Form[T]
// binds the request body into a *LoginRequest, runs Rules(), flashes
// errors via FlashErrors/FlashInput, redirects back, and returns
// router.ErrValidationAborted to short-circuit the handler. The frontend
// reads the flashed errors/old props on the next render.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func (r *LoginRequest) Rules() validation.Rules {
	return validation.Rules{
		"email":    {validation.Required(), validation.Email()},
		"password": {validation.Required()},
	}
}

// RegisterRequest schema for POST /register. The validation.Unique rule
// is enforced by the validation engine against the configured database
// driver; validation.Confirmed pairs password with password_confirmation
// without a manual equality check.
type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r *RegisterRequest) Rules() validation.Rules {
	return validation.Rules{
		"name":     {validation.Required(), validation.Max(255)},
		"email":    {validation.Required(), validation.Email(), validation.Unique("users", "email")},
		"password": {validation.Required(), validation.Min(8), validation.Confirmed()},
	}
}

// AuthShowLoginForm displays the login page
func AuthShowLoginForm(ctx *router.Context) error {
	return view.Render(ctx, "Auth/Login", view.Props{})
}

// AuthLogin handles the login request
func AuthLogin(ctx *router.Context) error {
	req, err := vform.Form[LoginRequest](ctx)
	if err != nil {
		return err
	}
	manager, err := ctx.Auth()
	if err != nil {
		return err
	}

	credentials := map[string]interface{}{
		"email":    req.Email,
		"password": req.Password,
	}

	success, _ := manager.Attempt(ctx.Response, ctx.Request, credentials, req.Remember)
	if !success {
		ctx.FlashErrors(map[string][]string{
			"email": {"These credentials do not match our records."},
		})
		ctx.FlashInput(map[string]any{"email": req.Email})
		return view.Back(ctx)
	}

	// Honour the intended destination the auth middleware stashed in the
	// session before bouncing the guest to /login (falls back to
	// /dashboard for a direct login).
	return view.Redirect(ctx, ctx.Intended("/dashboard"))
}

// AuthLogout handles the logout request
func AuthLogout(ctx *router.Context) error {
	manager, err := ctx.Auth()
	if err != nil {
		return err
	}
	if err := manager.Logout(ctx.Response, ctx.Request); err != nil {
		return err
	}
	return view.Redirect(ctx, "/login")
}

// AuthShowRegisterForm displays the registration page
func AuthShowRegisterForm(ctx *router.Context) error {
	return view.Render(ctx, "Auth/Register", view.Props{})
}

// AuthRegister handles the registration request
func AuthRegister(ctx *router.Context) error {
	req, err := vform.Form[RegisterRequest](ctx)
	if err != nil {
		return err
	}
	manager, err := ctx.Auth()
	if err != nil {
		return err
	}

	hashedPassword, err := manager.Hash(req.Password)
	if err != nil {
		ctx.Log().Error("Failed to hash password", "error", err)
		ctx.FlashErrors(map[string][]string{"password": {"Failed to process password."}})
		ctx.FlashInput(map[string]any{"name": req.Name, "email": req.Email})
		return view.Back(ctx)
	}

	user, err := models.User{}.Create(ctx.Request.Context(), map[string]any{
		"name":     req.Name,
		"email":    req.Email,
		"password": hashedPassword,
	})
	if err != nil {
		ctx.Log().Error("Failed to create user", "error", err)
		ctx.FlashErrors(map[string][]string{"email": {"Failed to create account. Please try again."}})
		ctx.FlashInput(map[string]any{"name": req.Name, "email": req.Email})
		return view.Back(ctx)
	}

	ctx.Log().Info("User created successfully", "email", user.Email, "id", user.ID)

	credentials := map[string]interface{}{
		"email":    req.Email,
		"password": req.Password,
	}
	if success, _ := manager.Attempt(ctx.Response, ctx.Request, credentials, false); success {
		return view.Redirect(ctx, ctx.Intended("/dashboard"))
	}
	return view.Redirect(ctx, "/login")
}
