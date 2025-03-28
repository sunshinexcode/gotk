package vmetric_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestInit(t *testing.T) {
	metric, err := vmetric.Init(vmap.M{"DisableGoCollector": true, "DisableProcessCollector": true})

	vtest.Nil(t, err)
	vtest.Equal(t, "", metric.Options.Namespace)
	vtest.Equal(t, true, metric.Options.DisableGoCollector)
	vtest.Equal(t, true, metric.Options.DisableProcessCollector)
}

func TestInitError(t *testing.T) {
	metric, err := vmetric.Init(vmap.M{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
	vtest.Equal(t, "", metric.Options.Namespace)
}

func TestNew(t *testing.T) {
	metric, err := vmetric.New(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, "", metric.Options.Namespace)
}

func TestD(t *testing.T) {
	metric, err := vmetric.New(nil)

	vtest.Nil(t, err)
	vtest.Nil(t, metric.D())
}

func TestRunMock(t *testing.T) {
	metric, err := vmetric.New(nil)

	patchHandle := vmock.ApplyFuncReturn(http.Handle)
	defer patchHandle.Reset()

	patchListenAndServe := vmock.ApplyFuncReturn(http.ListenAndServe, nil)
	defer patchListenAndServe.Reset()

	vtest.Nil(t, err)
	vtest.Nil(t, metric.Run())
}

func TestRunErrorMock(t *testing.T) {
	patchHandle := vmock.ApplyFuncReturn(http.Handle)
	defer patchHandle.Reset()

	patchListenAndServe := vmock.ApplyFuncReturn(http.ListenAndServe, errors.New("server error"))
	defer patchListenAndServe.Reset()

	patchFatal := vmock.ApplyFuncReturn(vlog.Fatal)
	defer patchFatal.Reset()

	patchNew := vmock.ApplyFuncReturn(vmetric.New, &vmetric.Metric{Options: &vmetric.Options{Url: "/", Port: ":9091"}}, nil)
	defer patchNew.Reset()

	metric, err := vmetric.New(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, "/", metric.Options.Url)
	vtest.Nil(t, metric.Run())

	time.Sleep(100 * time.Millisecond)
}

func TestSetConfig(t *testing.T) {
	metric, err := vmetric.New(nil)

	vtest.Nil(t, err)
	vtest.Nil(t, metric.SetConfig(nil))
}

func TestT(t *testing.T) {
	metric, err := vmetric.New(nil)

	vtest.Nil(t, err)
	vtest.Nil(t, metric.T())
}
