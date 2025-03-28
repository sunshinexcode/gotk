package vmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestElapsedMiddleware(t *testing.T) {
	metric, err := vmetric.Init(vmap.M{"DisableGoCollector": true, "DisableProcessCollector": true, "HttpRequestDurationLabelNames": []string{"type", "api", "code"}, "HttpRequestTotalLabelNames": []string{"type", "api", "code"}})
	vtest.Nil(t, err)

	defer vmetric.Reset()

	g := vapi.Default()
	g.GET("/test1", vmiddleware.ElapsedMiddleware(metric), func(c *vapi.Context) {
		voutput.O(c, nil, "test1")
	})
	g.Any("/test2", vmiddleware.ElapsedMiddleware(metric), func(c *vapi.Context) {
		time.Sleep(200 * time.Millisecond)
		voutput.O(c, nil, "test2")
	})

	req := httptest.NewRequest(http.MethodGet, "/test1", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test1", vjson.Parse(w.Body.String()).Get("data").Str)

	req = httptest.NewRequest(http.MethodPost, "/test2", nil)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test2", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestElapsedBusinessMiddleware(t *testing.T) {
	metric, err := vmetric.Init(vmap.M{"DisableGoCollector": true, "DisableProcessCollector": true, "HttpRequestDurationLabelNames": []string{"type", "api", "code", "business"}, "HttpRequestTotalLabelNames": []string{"type", "api", "code", "business"}})
	vtest.Nil(t, err)

	defer vmetric.Reset()

	g := vapi.Default()
	g.GET("/test1", vmiddleware.BasicAuthMiddleware(vapi.Accounts{"test": "test"}), vmiddleware.ParseBasicAuthMiddleware("business"), vmiddleware.ElapsedBusinessMiddleware(metric), func(c *vapi.Context) {
		voutput.O(c, nil, "test1")
	})
	g.Any("/test2", vmiddleware.BasicAuthMiddleware(vapi.Accounts{"test": "test"}), vmiddleware.ParseBasicAuthMiddleware("business"), vmiddleware.ElapsedBusinessMiddleware(metric), func(c *vapi.Context) {
		time.Sleep(200 * time.Millisecond)
		voutput.O(c, nil, "test2")
	})

	req := httptest.NewRequest(http.MethodGet, "/test1", nil)
	req.Header.Add("Authorization", "Basic dGVzdDp0ZXN0")
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test1", vjson.Parse(w.Body.String()).Get("data").Str)

	req = httptest.NewRequest(http.MethodPost, "/test2", nil)
	req.Header.Add("Authorization", "Basic dGVzdDp0ZXN0")
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test2", vjson.Parse(w.Body.String()).Get("data").Str)
}
