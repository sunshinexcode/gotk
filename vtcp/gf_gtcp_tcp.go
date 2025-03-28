package vtcp

import (
	"crypto/tls"
	"time"

	"github.com/gogf/gf/v2/net/gtcp"
)

type (
	Conn       = gtcp.Conn
	PkgOption  = gtcp.PkgOption
	ServerGtcp = gtcp.Server
)

func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error) {
	return gtcp.LoadKeyCrt(crtFile, keyFile)
}

func NewConn(addr string, timeout ...time.Duration) (*Conn, error) {
	return gtcp.NewConn(addr, timeout...)
}

func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error) {
	return gtcp.NewConnTLS(addr, tlsConfig)
}

func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*Conn), name ...string) (*ServerGtcp, error) {
	return gtcp.NewServerKeyCrt(address, crtFile, keyFile, handler, name...)
}
