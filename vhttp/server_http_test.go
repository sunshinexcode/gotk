package vhttp_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fvbock/endless"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewServer(t *testing.T) {
	s := &vhttp.Server{Addr: ":8000", G: vapi.Default(), StopHandler: func() {}}
	testController := func(e *vapi.Engine) {
		e.GET("/test", func(c *vapi.Context) {
			voutput.O(c, nil, "test")
		})
	}
	app := vfx.New(
		vfx.Options(
			vfx.Supply(s),
			vfx.Provide(vhttp.NewServer),
			vfx.Invoke(testController),
		))
	ctx := context.TODO()

	vtest.Nil(t, app.Start(ctx))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	s.G.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)

	dataJson := vjson.Parse(w.Body.String())

	vtest.Equal(t, int64(0), dataJson.Get("code").Int())
	vtest.Equal(t, "success", dataJson.Get("msg").Str)
	vtest.Equal(t, "test", dataJson.Get("data").Str)

	vtest.Nil(t, app.Stop(ctx))
}

func TestNewServerErrorMock(t *testing.T) {
	patchListenAndServe := vmock.ApplyFuncReturn(endless.ListenAndServe, errors.New("server error"))
	defer patchListenAndServe.Reset()

	patchFatal := vmock.ApplyFuncReturn(vlog.Fatal)
	defer patchFatal.Reset()

	s := &vhttp.Server{Addr: ":8000", G: vapi.Default(), StopHandler: func() {}}
	testController := func(e *vapi.Engine) {
		e.GET("/test", func(c *vapi.Context) {
			voutput.O(c, nil, "test")
		})
	}
	app := vfx.New(
		vfx.Options(
			vfx.Supply(s),
			vfx.Provide(vhttp.NewServer),
			vfx.Invoke(testController),
		))
	ctx := context.TODO()

	vtest.Nil(t, app.Start(ctx))
	vtest.Nil(t, app.Stop(ctx))
}
