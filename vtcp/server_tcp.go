package vtcp

import (
	"context"

	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
)

type IHandler interface {
	Handle(conn *Conn)
}

type Server struct {
	Address     string `mapstructure:",omitempty"`
	CrtFile     string `mapstructure:",omitempty"`
	Handler     func(conn *Conn)
	KeyFile     string `mapstructure:",omitempty"`
	StopHandler func()
}

// NewServer new server
func NewServer(lifecycle vfx.Lifecycle, s *Server) *Server {
	return s.NewServer(lifecycle)
}

// NewServer new server
func (s *Server) NewServer(lifecycle vfx.Lifecycle) *Server {
	lifecycle.Append(vfx.Hook{
		OnStart: func(ctx context.Context) error {
			vlog.Info("start tcp server")

			go func() {
				server, err := NewServerKeyCrt(s.Address, s.CrtFile, s.KeyFile, func(conn *Conn) {
					s.Handler(conn)
				})

				if err != nil {
					vlog.Fatal("start tcp server", "err", err)
					return
				}

				_ = server.Run()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			vlog.Info("stop tcp server")

			s.StopHandler()

			return nil
		},
	})

	return s
}
