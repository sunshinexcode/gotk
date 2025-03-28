package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vwebsocket"
)

// NewWebsocketServer new websocket server
func NewWebsocketServer(lifecycle vfx.Lifecycle, config vconfig.IConfig, handler vwebsocket.IHandler) *vwebsocket.Server {
	s := &vwebsocket.Server{Port: config.GetWebsocketServer().Port, Handlers: handler.GetHandlers(), StopHandler: func() {}}

	return vwebsocket.NewServer(lifecycle, s)
}
