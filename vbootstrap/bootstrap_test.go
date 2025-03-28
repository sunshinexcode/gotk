package vbootstrap_test

import (
	"go.uber.org/fx"

	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vwebsocket"
)

type Handler struct {
}

func (h *Handler) Handle(conn *vtcp.Conn) {
}

func (h *Handler) GetHandlers() []vwebsocket.Handler {
	return nil
}

type lifecycleWrapper struct {
}

func (l *lifecycleWrapper) Append(h fx.Hook) {
}
