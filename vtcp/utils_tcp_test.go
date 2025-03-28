package vtcp_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCloseConnMock(t *testing.T) {
	conn := vtcp.Conn{Conn: &vtcp.ConnMock{}}

	vtest.Nil(t, vtcp.CloseConn(&conn))

	conn2 := &vtcp.Conn{}
	patch := vmock.ApplyMethodReturn(conn2, "Close", nil)
	defer vmock.Reset(patch)

	vtest.Nil(t, vtcp.CloseConn(&vtcp.Conn{Conn: conn2}))
}

func TestCloseConnErrorMock(t *testing.T) {
	conn := vtcp.Conn{Conn: &vtcp.ConnCloseErrMock{}}

	err := vtcp.CloseConn(&conn)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10160|close connection error|<nil> \n-> close error", err.Error())

	conn2 := &vtcp.Conn{}
	patch := vmock.ApplyMethodReturn(conn2, "Close", errors.New("close error"))
	defer vmock.Reset(patch)

	err = vtcp.CloseConn(&vtcp.Conn{Conn: conn2})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10160|close connection error|<nil> \n-> close error", err.Error())
}
