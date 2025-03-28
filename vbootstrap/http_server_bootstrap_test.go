package vbootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewHttpServer(t *testing.T) {
	e := vbootstrap.NewHttpServer(&lifecycleWrapper{}, &vconfig.Config{})

	vtest.Equal(t, "/", e.BasePath())
}
