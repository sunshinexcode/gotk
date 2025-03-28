package vcontroller

import (
	"github.com/gin-gonic/gin/binding"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/voutput"
)

type IBaseController interface {
	Route()
}

type BaseControllerParam struct {
	vfx.In

	Metric *vmetric.Metric
}

type BaseController struct {
	Metric *vmetric.Metric
}

func NewBaseController(p BaseControllerParam) *BaseController {
	return &BaseController{Metric: p.Metric}
}

func (controller *BaseController) BindBody(c *vapi.Context, req any) (err error) {
	if err = c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		voutput.O(c, verror.Wrap(err, vcode.CodeErrParamBind, req), nil)
	}

	return
}
