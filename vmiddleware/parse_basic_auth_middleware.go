package vmiddleware

import (
	"strings"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vauth"
	"github.com/sunshinexcode/gotk/vbase64"
)

func ParseBasicAuthMiddleware(userNameKey string) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		authorization := c.GetHeader(vauth.BasicAuthorizationKey)

		if authorization != "" {
			authorizationList := strings.SplitN(authorization, " ", 2)

			if len(authorizationList) == 2 && authorizationList[0] == "Basic" {
				if auth, err := vbase64.DecodeToStr(authorizationList[1]); err == nil {
					authList := strings.SplitN(auth, ":", 2)
					if len(authList) == 2 {
						c.Set(userNameKey, authList[0])
						c.Set(vauth.BasicAuthorizationKey, authorization)
					}
				}
			}
		}

		c.Next()
	}
}
