package vcode_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCodeNewCode(t *testing.T) {
	c := vcode.NewCode(1, "test", nil)

	vtest.Equal(t, 1, c.Code())
	vtest.Equal(t, "1", c.CodeStr())
	vtest.Equal(t, "test", c.Message())
	vtest.Equal(t, nil, c.Data())
}

func TestCodeSetData(t *testing.T) {
	c := vcode.NewCode(1, "test", nil)

	vtest.Equal(t, nil, c.Data())

	c.SetData("test")

	vtest.Equal(t, "test", c.Data())
}

func TestCodeSetMessage(t *testing.T) {
	c := vcode.NewCode(1, "test", nil)

	vtest.Equal(t, "test", c.Message())

	c.SetMessage("invalid")

	vtest.Equal(t, 1, c.Code())
	vtest.Equal(t, "invalid", c.Message())
}
