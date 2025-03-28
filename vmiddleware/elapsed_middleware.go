package vmiddleware

import (
	"time"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vauth"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmetric"
)

// ElapsedMiddleware elapsed time
func ElapsedMiddleware(metric *vmetric.Metric) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		timeStart := time.Now()

		defer func() {
			lostTime := time.Since(timeStart)
			code := c.GetInt("code")
			codeStr := vconv.String(code)

			vlog.Infoc(c, "request api", "api", c.FullPath(), "code", code, "clientIp", c.ClientIP(), "lostTimeMs", lostTime.Milliseconds())

			vmetric.MetricHttpRequestTotalTypeApi(metric, c.FullPath(), codeStr)
			vmetric.MetricHttpRequestDurationTypeApi(metric, timeStart, c.FullPath(), codeStr)
		}()

		c.Next()
	}
}

// ElapsedBusinessMiddleware elapsed time
func ElapsedBusinessMiddleware(metric *vmetric.Metric) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		timeStart := time.Now()

		defer func() {
			lostTime := time.Since(timeStart)
			code := c.GetInt("code")
			codeStr := vconv.String(code)
			business := vauth.GetBasicAuthUserName(c)

			vlog.Infoc(c, "request api", "api", c.FullPath(), "code", code, "business", business, "clientIp", c.ClientIP(), "lostTimeMs", lostTime.Milliseconds())

			vmetric.MetricHttpRequestTotalTypeApi(metric, c.FullPath(), codeStr, business)
			vmetric.MetricHttpRequestDurationTypeApi(metric, timeStart, c.FullPath(), codeStr, business)
		}()

		c.Next()
	}
}
