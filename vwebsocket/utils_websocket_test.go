package vwebsocket_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vwebsocket"
)

func TestCloseConnMock(t *testing.T) {
	conn := vwebsocket.WebSocket{}

	patch := vmock.ApplyMethodReturn(&vwebsocket.Conn{}, "Close", nil)
	defer vmock.Reset(patch)

	vtest.Nil(t, vwebsocket.CloseConn(&conn))
}

func TestCloseConnErrorMock(t *testing.T) {
	conn := vwebsocket.WebSocket{}

	patch := vmock.ApplyMethodReturn(&vwebsocket.Conn{}, "Close", errors.New("close error"))
	defer vmock.Reset(patch)

	err := vwebsocket.CloseConn(&conn)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10160|close connection error|<nil> \n-> close error", err.Error())
}
