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

func TestBasicAuthMiddleware(t *testing.T) {
	g := vapi.Default()
	g.GET("/test", vmiddleware.BasicAuthMiddleware(vapi.Accounts{"test": "9WIiOZq6bE2UgJzX"}), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Authorization", "Basic dGVzdDo5V0lpT1pxNmJFMlVnSnpY")
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestBasicAuthMiddlewareError(t *testing.T) {
	g := vapi.Default()
	g.GET("/test", vmiddleware.BasicAuthMiddleware(vapi.Accounts{"test": "9WIiOZq6bE2UgJzX"}), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusUnauthorized, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "", vjson.Parse(w.Body.String()).Get("data").Str)
}
