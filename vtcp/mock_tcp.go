package vtcp

import (
	"errors"
	"net"
	"time"
)

type AddrMock struct {
}

func (a *AddrMock) Network() string {
	return ""
}

func (a *AddrMock) String() string {
	return ""
}

type ConnMock struct {
}

func (c *ConnMock) Close() error {
	return nil
}

func (c *ConnMock) LocalAddr() net.Addr {
	return nil
}

func (c *ConnMock) Read(b []byte) (n int, err error) {
	return
}

func (c *ConnMock) RemoteAddr() net.Addr {
	return nil
}

func (c *ConnMock) SetDeadline(t time.Time) error {
	return nil
}

func (c *ConnMock) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *ConnMock) SetWriteDeadline(t time.Time) error {
	return nil
}

func (c *ConnMock) Write(b []byte) (n int, err error) {
	return
}

type ConnCloseErrMock struct {
	ConnMock
}

func (c *ConnCloseErrMock) Close() error {
	return errors.New("close error")
}
