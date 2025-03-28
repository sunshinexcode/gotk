package vpprof

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
)

func New(lc vfx.Lifecycle) {
	lc.Append(vfx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				runtime.SetMutexProfileFraction(1)
				runtime.SetBlockProfileRate(1)

				if err := http.ListenAndServe(":6060", nil); err != nil {
					vlog.Error("ListenAndServe", "err", err)
				}
			}()

			return nil
		},
	})
}
