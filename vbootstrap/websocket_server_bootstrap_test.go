package vbootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewWebsocketServer(t *testing.T) {
	s := vbootstrap.NewWebsocketServer(&lifecycleWrapper{}, &vconfig.Config{}, &Handler{})

	vtest.Equal(t, int(0), s.Port)
}
