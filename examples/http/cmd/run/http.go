package run

import (
	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vversion"

	"app/configs"
	"app/internal/bootstrap"
	"app/internal/cache"
	"app/internal/controller"
	"app/internal/model"
	"app/internal/service"
	"app/internal/thirdparty"
)

var (
	httpCmd = &vcmd.Command{
		Use:   "http",
		Short: "Start http REST API",
		Run:   initHttp,
	}

	versionCmd = &vcmd.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *vcmd.Command, args []string) {
			vversion.Print()
		},
	}
)

func initHttp(cmd *vcmd.Command, args []string) {
	vfx.New(
		// vfx.NopLogger,
		vfx.Options(
			vfx.Provide(
				configs.NewConfig,
				bootstrap.NewHttpServer,
			),
			vfx.Invoke(
				bootstrap.NewLog,
			),
			controller.Module,
			service.Module,
			model.Module,
			cache.Module,
			thirdparty.Module,
			bootstrap.Module,
		)).Run()
}
