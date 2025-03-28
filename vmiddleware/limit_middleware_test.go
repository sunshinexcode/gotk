package vmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vlimit"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestLimitMiddleware(t *testing.T) {
	g := vapi.Default()
	g.GET("/test", vmiddleware.LimitMiddleware(vlimit.New(0.1, 1)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		g.ServeHTTP(w, req)

		vtest.Equal(t, http.StatusOK, w.Code)
		vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
		vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
		vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
	}()

	go func() {
		time.Sleep(1 * time.Millisecond)
		defer wg.Done()

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		g.ServeHTTP(w, req)

		vtest.Equal(t, http.StatusTooManyRequests, w.Code)
		vtest.Equal(t, int64(vcode.CodeErrRateLimit.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
		vtest.Equal(t, vcode.CodeErrRateLimit.Message(), vjson.Parse(w.Body.String()).Get("msg").Str)
	}()

	wg.Wait()
}
