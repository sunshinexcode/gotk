package vwebsocket

import (
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
)

func CloseConn(conn *WebSocket) (err error) {
	if err = conn.Close(); err != nil {
		return verror.Wrap(err, vcode.CodeErrCloseConn)
	}

	return
}
