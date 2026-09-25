package app

import (
	"github.com/velocitykode/velocity/contract"
)

// Errors adds the application's own rules to the error handler. main.go
// passes it to v.Errors(...). The framework already maps its own errors
// (unmatched routes 404/405, validation 422, unauthenticated 401 or a login
// redirect, forbidden 403, CSRF 419, ORM not-found 404, panics 500) and
// renders application/problem+json for JSON clients and the VIEW_ERROR_PAGE
// component for browsers, so this stays empty until the app has errors of
// its own.
func Errors(h contract.ErrorHandler) {
	// Answer an app sentinel with a 410 (imports velocity/problem):
	//
	//	problem.MapIs(h, ErrInviteExpired, func(err error) error {
	//		return problem.Gone("This invite has expired.").WithCause(err)
	//	})
	//
	// Write the answer for an app error type yourself (imports
	// velocity/problem and net/http); returning false falls through to the
	// negotiated problem+json or error page:
	//
	//	problem.RenderFor[*BillingError](h, func(rc problem.RenderContext, err *BillingError, _ *problem.ErrorContext) bool {
	//		if rc.WantsJSON() {
	//			return false
	//		}
	//		return rc.Redirect(http.StatusSeeOther, "/billing") == nil
	//	})
}
