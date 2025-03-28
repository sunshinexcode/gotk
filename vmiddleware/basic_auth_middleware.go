package vmiddleware

import "github.com/sunshinexcode/gotk/vapi"

// BasicAuthMiddleware basic auth middleware
func BasicAuthMiddleware(accounts vapi.Accounts) vapi.HandlerFunc {
	return vapi.BasicAuth(accounts)
}
