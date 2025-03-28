package vmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCorsMiddleware(t *testing.T) {
	g := vapi.Default()
	g.GET("/test", vmiddleware.CorsMiddleware(), func(c *vapi.Context) {
		vtest.Equal(t, "*", c.Writer.Header().Get("Access-Control-Allow-Origin"))

		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)

}

func TestCorsMiddlewareMethodOptions(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.CorsMiddleware(), func(c *vapi.Context) {
		vtest.Equal(t, "*", c.Writer.Header().Get("Access-Control-Allow-Origin"))

		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusNoContent, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "", vjson.Parse(w.Body.String()).Get("msg").Str)
}
