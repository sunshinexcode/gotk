package vmiddleware_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtime"
)

func TestExpireMiddleware(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.ExpiredMiddleware[RequestExpire](vmiddleware.DefaultExpired), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodGet, vstr.S("/test?ts=%d", vtime.Timestamp()-1), nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)

	req = httptest.NewRequest(http.MethodGet, vstr.S("/test?ts=%d", vtime.Timestamp()-11*60), nil)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusRequestTimeout, w.Code)
	vtest.Equal(t, int64(vcode.CodeErrRequestExpired.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, vcode.CodeErrRequestExpired.Message(), vjson.Parse(w.Body.String()).Get("msg").Str)

	req = httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(fmt.Sprintf("ts=%d", vtime.Timestamp()))))
	req.Header.Add("Content-Type", vapi.MimePostForm)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)

	dataStr, err := vjson.Encode(vconv.Map(RequestExpire{Name: "test1", Ts: vtime.Timestamp() - 1}))

	vtest.Nil(t, err)

	req = httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestExpireMiddlewareSetExpiredKey(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.ExpiredMiddleware[RequestExpire](vmiddleware.NewExpired().
		SetExpiredKey("expired").SetGetExpiredFunc(vmiddleware.GetExpired)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	dataStr, err := vjson.Encode(vconv.Map(RequestExpire{Name: "test", Expired: vtime.Timestamp() - 1}))

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}
