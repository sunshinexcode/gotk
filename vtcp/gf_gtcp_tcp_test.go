package vtcp_test

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/gogf/gf/v2/net/gtcp"

	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewConnMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(gtcp.NewConn, &vtcp.Conn{}, nil)
	defer vmock.Reset(patch)

	_, err := vtcp.NewConn(":6000", 5*time.Second)

	vtest.Nil(t, err)
}

func TestNewConnTLSMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(gtcp.NewConnTLS, &vtcp.Conn{}, nil)
	defer vmock.Reset(patch)

	_, err := vtcp.NewConnTLS(":6000", &tls.Config{})

	vtest.Nil(t, err)
}

func TestNewServerKeyCrtMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(gtcp.NewServerKeyCrt, &vtcp.ServerGtcp{}, nil)
	defer vmock.Reset(patch)

	_, err := vtcp.NewServerKeyCrt(":6000", "", "", func(conn *vtcp.Conn) {}, "test")

	vtest.Nil(t, err)
}
