package controller

import (
	"github.com/sunshinexcode/gotk/vcontroller"
	"github.com/sunshinexcode/gotk/vfx"
)

var Module = vfx.Options(
	vfx.Provide(vcontroller.NewBaseController),
	vfx.Invoke(
		NewCompanyController,
		vcontroller.NewHealthController,
	),
)
