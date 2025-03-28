package vfx_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNew(t *testing.T) {
	vtest.Nil(t, vfx.New().Err())
}

func TestInvoke(t *testing.T) {
	vtest.Equal(t, "fx.Invoke()", vfx.Invoke().String())
}

func TestOptions(t *testing.T) {
	vtest.Equal(t, "fx.Options()", vfx.Options().String())
}

func TestProvide(t *testing.T) {
	vtest.Equal(t, "fx.Provide()", vfx.Provide().String())
}

func TestSupply(t *testing.T) {
	vtest.Equal(t, "fx.Supply()", vfx.Supply().String())
}
