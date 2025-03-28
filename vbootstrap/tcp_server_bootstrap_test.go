package vbootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewTcpServer(t *testing.T) {
	s := vbootstrap.NewTcpServer(&lifecycleWrapper{}, &vconfig.Config{}, &Handler{})

	vtest.Equal(t, "", s.Address)
}
