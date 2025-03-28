package vwebsocket_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vwebsocket"
)

func TestNewClient(t *testing.T) {
	vtest.Equal(t, 0, vwebsocket.NewClient().ReadBufferSize)
}
