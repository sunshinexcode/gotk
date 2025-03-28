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

func TestParseBasicAuthMiddleware(t *testing.T) {
	g := vapi.Default()
	g.GET("/test1", vmiddleware.ParseBasicAuthMiddleware("user"), func(c *vapi.Context) {
		userName, _ := c.Get("user")
		voutput.O(c, nil, userName)
	})
	g.Any("/test2", vmiddleware.ParseBasicAuthMiddleware("business"), func(c *vapi.Context) {
		userName, _ := c.Get("business")
		voutput.O(c, nil, userName)
	})

	req := httptest.NewRequest(http.MethodGet, "/test1", nil)
	req.Header.Add("Authorization", "Basic dGVzdDp0ZXN0")
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)

	req = httptest.NewRequest(http.MethodPost, "/test2", nil)
	req.Header.Add("Auth", "Bearer dGVzdDp0ZXN0")
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "", vjson.Parse(w.Body.String()).Get("data").Str)
}
