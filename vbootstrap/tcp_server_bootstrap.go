package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vtcp"
)

// NewTcpServer new tcp server
func NewTcpServer(lifecycle vfx.Lifecycle, config vconfig.IConfig, handler vtcp.IHandler) *vtcp.Server {
	s := &vtcp.Server{Address: config.GetTcpServer().Address, CrtFile: config.GetTcpServer().CrtFile, Handler: handler.Handle,
		KeyFile: config.GetTcpServer().KeyFile, StopHandler: func() {}}

	return vtcp.NewServer(lifecycle, s)
}
