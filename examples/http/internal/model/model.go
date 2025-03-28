package model

import (
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vmodel"
)

var Module = vfx.Provide(
	vmodel.NewBaseModel,
	NewCompanyModel,
)
