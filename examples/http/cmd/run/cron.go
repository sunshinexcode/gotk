package run

import (
	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vfx"

	"app/configs"
	"app/internal/bootstrap"
	"app/internal/cache"
	"app/internal/model"
	"app/internal/service"
	"app/internal/thirdparty"
)

var (
	cronCmd = &vcmd.Command{
		Use:   "cron",
		Short: "Start cron",
		Run:   initCron,
	}
)

func initCron(cmd *vcmd.Command, args []string) {
	vfx.New(
		// vfx.NopLogger,
		vfx.Options(
			vfx.Provide(
				configs.NewConfig,
			),
			vfx.Invoke(
				bootstrap.NewLog,
				createTaskForVerify,
				processTaskForVerify,
			),
			service.Module,
			model.Module,
			cache.Module,
			thirdparty.Module,
			bootstrap.Module,
		)).Run()
}

func createTaskForVerify(iCompanyCronService service.ICompanyCronService) {
	iCompanyCronService.CronCreateTaskForVerify()
}

func processTaskForVerify(iCompanyCronService service.ICompanyCronService) {
	iCompanyCronService.CronProcessTaskForVerify()
}
