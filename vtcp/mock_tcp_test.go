package vtcp_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtime"
)

func initAddrMock() (addr *vtcp.AddrMock) {
	addr = &vtcp.AddrMock{}

	return
}

func initConnMock() (conn *vtcp.ConnMock) {
	conn = &vtcp.ConnMock{}

	return
}

func initConnCloseErrMock() (conn *vtcp.ConnCloseErrMock) {
	conn = &vtcp.ConnCloseErrMock{}

	return
}

func TestAddrMockNetwork(t *testing.T) {
	addr := initAddrMock()

	vtest.Equal(t, "", addr.Network())
}

func TestAddrMockString(t *testing.T) {
	addr := initAddrMock()

	vtest.Equal(t, "", addr.String())
}

func TestConnMockClose(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.Close())
}

func TestConnMockLocalAddr(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.LocalAddr())
}

func TestConnMockRead(t *testing.T) {
	conn := initConnMock()
	n, err := conn.Read(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, 0, n)
}

func TestConnMockRemoteAddr(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.RemoteAddr())
}

func TestConnMockSetDeadline(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.SetDeadline(vtime.GetNow()))
}

func TestConnMockSetReadDeadline(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.SetReadDeadline(vtime.GetNow()))
}

func TestConnMockSetWriteDeadline(t *testing.T) {
	conn := initConnMock()

	vtest.Nil(t, conn.SetWriteDeadline(vtime.GetNow()))
}

func TestConnMockWrite(t *testing.T) {
	conn := initConnMock()
	n, err := conn.Write(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, 0, n)
}

func TestConnCloseErrMockClose(t *testing.T) {
	conn := initConnCloseErrMock()

	vtest.NotNil(t, conn.Close())
	vtest.Equal(t, "close error", conn.Close().Error())
}
