package vwebsocket

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
)

type IHandler interface {
	GetHandlers() []Handler
}

type Server struct {
	EnableHttps bool   `mapstructure:",omitempty"`
	Port        int    `mapstructure:",omitempty"`
	CrtFile     string `mapstructure:",omitempty"`
	Handlers    []Handler
	KeyFile     string `mapstructure:",omitempty"`
	StopHandler func()
}

type Handler struct {
	HandlerFunc func(r *Request)
	Pattern     string `mapstructure:",omitempty"`
}

// NewServer new server
func NewServer(lifecycle vfx.Lifecycle, s *Server) *Server {
	return s.NewServer(lifecycle)
}

// NewServer new server
func (s *Server) NewServer(lifecycle vfx.Lifecycle) *Server {
	var svr *ServerGhttp

	lifecycle.Append(vfx.Hook{
		OnStart: func(ctx context.Context) error {
			vlog.Info("start websocket server")

			go func() {
				svr = g.Server()
				for _, handlers := range s.Handlers {
					svr.BindHandler(handlers.Pattern, handlers.HandlerFunc)
				}

				if s.EnableHttps {
					svr.EnableHTTPS(s.CrtFile, s.KeyFile)
				}

				svr.SetPort(s.Port)
				svr.Run()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			vlog.Info("stop websocket server")

			if err := svr.Shutdown(); err != nil {
				vlog.Error("Shutdown", "err", err)
			}

			s.StopHandler()

			return nil
		},
	})

	return s
}
