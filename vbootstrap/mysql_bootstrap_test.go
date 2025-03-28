package vbootstrap_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewMysqlErrorDecodeMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmap.Decode, errors.New("decode error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewMysql(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "decode error", err.Error())
}

func TestNewMysqlErrorNewMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vmysql.New, &vmysql.Mysql{}, errors.New("new error"))
	defer vmock.Reset(patch)

	_, err := vbootstrap.NewMysql(&vconfig.Config{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "new error", err.Error())
}
