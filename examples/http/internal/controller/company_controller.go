package controller

import (
	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcontroller"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vreq"

	"app/configs"
	"app/internal/req"
	"app/internal/service"
)

var _ vcontroller.IBaseController = (*CompanyController)(nil)

type CompanyControllerParam struct {
	vfx.In

	Engine         *vapi.Engine
	BaseController *vcontroller.BaseController

	ICompanyService service.ICompanyService
}

type CompanyController struct {
	Engine         *vapi.Engine
	BaseController *vcontroller.BaseController

	ICompanyService service.ICompanyService
}

func NewCompanyController(p CompanyControllerParam) *CompanyController {
	controller := &CompanyController{Engine: p.Engine, BaseController: p.BaseController, ICompanyService: p.ICompanyService}
	controller.Route()
	return controller
}

func (controller *CompanyController) Route() {
	group := controller.Engine.Group("/v1/company", vmiddleware.BasicAuthMiddleware(configs.BasicAuth), vmiddleware.TraceIdMiddleware[vreq.TraceIdReq](vmiddleware.DefaultTraceId))
	group.POST("/query", vmiddleware.ValidMiddleware[req.CompanyQueryReq](), vmiddleware.ElapsedMiddleware(controller.BaseController.Metric), controller.Query)
}

func (controller *CompanyController) Query(c *vapi.Context) {
	var request req.CompanyQueryReq

	if controller.BaseController.BindBody(c, &request) != nil {
		return
	}

	data, err := controller.ICompanyService.QueryLocalCache(c, &request)
	voutput.O(c, err, data)
}
