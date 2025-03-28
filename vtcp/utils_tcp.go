package vtcp

import (
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
)

func CloseConn(conn *Conn) (err error) {
	if err = conn.Close(); err != nil {
		return verror.Wrap(err, vcode.CodeErrCloseConn)
	}

	return
}
