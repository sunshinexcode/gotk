package vbootstrap_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewMetricErrorDecodeMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmap.Decode, errors.New("decode error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewMetric(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "decode error", err.Error())
}

func TestNewMetricErrorInitMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmetric.Init, &vmetric.Metric{}, errors.New("init error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewMetric(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "init error", err.Error())
}

func TestNewMetricErrorRunMock(t *testing.T) {
	patchInit := vmock.ApplyFuncReturn(vmetric.Init, &vmetric.Metric{}, nil)
	defer patchInit.Reset()

	patch := vmock.ApplyMethodReturn(&vmetric.Metric{}, "Run", errors.New("run error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewMetric(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "run error", err.Error())
}
