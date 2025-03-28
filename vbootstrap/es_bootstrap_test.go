package vbootstrap_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/ves"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewEsErrorDecodeMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmap.Decode, errors.New("decode error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewEs(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "decode error", err.Error())
}

func TestNewEsErrorNewMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(ves.New, &ves.Es{}, errors.New("new error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewEs(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "new error", err.Error())
}
