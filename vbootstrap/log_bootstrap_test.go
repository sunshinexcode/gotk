package vbootstrap_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewLogErrorDecodeMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmap.Decode, errors.New("decode error"))
	defer vmock.Reset(patch)

	err := vbootstrap.NewLog(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "decode error", err.Error())
}

func TestNewLogErrorSetConfigMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vlog.SetConfig, nil, errors.New("set config error"))
	defer vmock.Reset(patch)

	err := vbootstrap.NewLog(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "set config error", err.Error())
}
