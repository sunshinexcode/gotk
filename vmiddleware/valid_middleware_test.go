package vmiddleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestValidMiddleware(t *testing.T) {
	g := vapi.Default()
	g.GET("/user", vmiddleware.ValidMiddleware[RequestValid](), func(c *vapi.Context) {
		voutput.O(c, nil, "")
	})

	req := httptest.NewRequest(http.MethodGet, "/user?age=18", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(vcode.CodeErrParamInvalid.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "name is required", vjson.Parse(w.Body.String()).Get("msg").Str)

	req = httptest.NewRequest(http.MethodGet, "/user?name=test", nil)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(vcode.CodeErrParamInvalid.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "The Age value `0` must be equal or greater than 18", vjson.Parse(w.Body.String()).Get("msg").Str)

	req = httptest.NewRequest(http.MethodGet, "/user?name=test&age=20", nil)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(vcode.CodeSuccess.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, vcode.CodeSuccess.Message(), vjson.Parse(w.Body.String()).Get("msg").Str)
}

func TestValidMiddlewarePost(t *testing.T) {
	g := vapi.Default()
	g.POST("/user", vmiddleware.ValidMiddleware[RequestValid](vcode.NewCode(100, "test", nil)), func(c *vapi.Context) {
		voutput.O(c, nil, "")
	})

	req := httptest.NewRequest(http.MethodPost, "/user?req=1&name=test", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(100), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "The Age value `0` must be equal or greater than 18", vjson.Parse(w.Body.String()).Get("msg").Str)

	dataStr, err := vjson.Encode(vmap.M{"name": "test"})

	vtest.Nil(t, err)

	req = httptest.NewRequest(http.MethodPost, "/user?req=2&age=18", bytes.NewBuffer([]byte(dataStr)))
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(100), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "name is required", vjson.Parse(w.Body.String()).Get("msg").Str)

	req = httptest.NewRequest(http.MethodPost, "/user?req=3&age=18", bytes.NewBuffer([]byte("name=test")))
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(100), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "name is required", vjson.Parse(w.Body.String()).Get("msg").Str)
}

func TestValidMiddlewarePostMimeJson(t *testing.T) {
	g := vapi.Default()
	g.POST("/user", vmiddleware.ValidMiddleware[RequestValid](vcode.NewCode(100, "test", nil)), func(c *vapi.Context) {
		voutput.O(c, nil, "")
	})

	dataStr, err := vjson.Encode(vmap.M{"name": "test"})

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/user?req=5", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(100), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "The Age value `0` must be equal or greater than 18", vjson.Parse(w.Body.String()).Get("msg").Str)
}

func TestValidMiddlewarePostMimePostForm(t *testing.T) {
	g := vapi.Default()
	g.POST("/user", vmiddleware.ValidMiddleware[RequestValid](vcode.NewCode(100, "test", nil)), func(c *vapi.Context) {
		voutput.O(c, nil, "")
	})

	req := httptest.NewRequest(http.MethodPost, "/user?req=4", bytes.NewBuffer([]byte("name=test")))
	req.Header.Add("Content-Type", vapi.MimePostForm)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(100), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "The Age value `0` must be equal or greater than 18", vjson.Parse(w.Body.String()).Get("msg").Str)
}
