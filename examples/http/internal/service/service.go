package service

import "github.com/sunshinexcode/gotk/vfx"

var Module = vfx.Provide(
	NewCompanyCronService,
	NewCompanyService,
)
