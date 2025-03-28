package bootstrap

import (
	"github.com/sunshinexcode/gotk/vfx"
)

var Module = vfx.Options(
	vfx.Provide(
		//NewEs,
		NewMetric,
		//NewMongodb,
		NewMysql,
		NewRedis,
	),
	//vpprof.Module,
)
