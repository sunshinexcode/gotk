package vhttp

import (
	"context"

	"github.com/fvbock/endless"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
)

type Server struct {
	Addr        string
	G           *vapi.Engine
	StopHandler func()
}

// NewServer new server
func NewServer(lifecycle vfx.Lifecycle, s *Server) *vapi.Engine {
	return s.NewServer(lifecycle)
}

// NewServer new server
func (s *Server) NewServer(lifecycle vfx.Lifecycle) *vapi.Engine {
	lifecycle.Append(vfx.Hook{
		OnStart: func(ctx context.Context) error {
			vlog.Info("start http server")

			go func() {
				if err := endless.ListenAndServe(s.Addr, s.G); err != nil {
					vlog.Fatal("start http server", "err", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			vlog.Info("stop http server")

			s.StopHandler()

			return nil
		},
	})

	return s.G
}
