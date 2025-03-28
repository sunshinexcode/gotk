package vcontroller

import (
	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/voutput"
)

var _ IBaseController = (*HealthController)(nil)

type HealthControllerParam struct {
	vfx.In

	Engine *vapi.Engine
}

type HealthController struct {
	Engine *vapi.Engine
}

func NewHealthController(p HealthControllerParam) *HealthController {
	controller := &HealthController{Engine: p.Engine}
	controller.Route()
	return controller
}

func (controller *HealthController) Route() {
	controller.Engine.GET("/health/check", controller.Check)
	controller.Engine.GET("/", controller.Check)
}

func (controller *HealthController) Check(c *vapi.Context) {
	voutput.O(c, nil, nil)
}
