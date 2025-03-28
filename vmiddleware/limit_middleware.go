package vmiddleware

import (
	"net/http"

	"golang.org/x/time/rate"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/voutput"
)

// LimitMiddleware rate limit
func LimitMiddleware(limiter *rate.Limiter) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		if !limiter.Allow() {
			voutput.E(c, verror.ErrRateLimit, http.StatusTooManyRequests).Abort()
			return
		}

		c.Next()
	}
}
