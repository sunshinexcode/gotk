package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcontroller"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vreq"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/controller"
	"app/internal/req"
	"app/internal/resp"
	"app/internal/service"
	"app/internal/utils"
)

const (
	CompanyControllerUrlQuery = "/v1/company/query"
)

func initCompanyController() (g *vapi.Engine, c *controller.CompanyController, patches []*vmock.Patches) {
	g = vapi.Default()

	patchElapsedMiddleware := vmock.ApplyFuncReturn(vmiddleware.ElapsedMiddleware, vapi.HandlerFunc(func(c *vapi.Context) {}))
	patches = append(patches, patchElapsedMiddleware)

	c = controller.NewCompanyController(controller.CompanyControllerParam{Engine: g, BaseController: &vcontroller.BaseController{}, ICompanyService: &service.CompanyService{}})
	return
}

func TestCompanyControllerQuery(t *testing.T) {
	if !configs.TestOpen {
		return
	}

	url, resp, err := utils.Request(CompanyControllerUrlQuery, req.CompanyQueryReq{TraceIdReq: vreq.TraceIdReq{TraceId: "test"}, Id: 116711})
	vlog.Debug("TestCompanyControllerQuery", "err", err, "url", url, "resp", resp)

	vtest.Nil(t, err)
	vtest.Equal(t, http.StatusOK, resp.StatusCode())
	vtest.Equal(t, int64(0), vjson.Parse(string(resp.Body())).Get("code").Int())
	vtest.Equal(t, true, vjson.Parse(string(resp.Body())).Get("data").Get("companyName").Str != "")
}

func TestCompanyControllerQueryMock(t *testing.T) {
	g, c, patches := initCompanyController()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(c.ICompanyService, "QueryLocalCache", &resp.CompanyQueryResp{Id: 1}, nil)
	defer vmock.Reset(patch)

	body, _ := json.Marshal(req.CompanyQueryReq{TraceIdReq: vreq.TraceIdReq{TraceId: "test"}, Id: 1})
	req := httptest.NewRequest(http.MethodPost, CompanyControllerUrlQuery, bytes.NewBuffer(body))
	utils.SetHeader(req)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, int64(1), vjson.Parse(w.Body.String()).Get("data").Get("id").Int())
}

func TestCompanyControllerQueryErrorAuth(t *testing.T) {
	g, _, patches := initCompanyController()
	defer vmock.ResetMock(patches)

	req := httptest.NewRequest(http.MethodPost, CompanyControllerUrlQuery, nil)
	req.Header.Set("Content-Type", vapi.MimeJson)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCompanyControllerQueryErrorParam(t *testing.T) {
	g, _, patches := initCompanyController()
	defer vmock.ResetMock(patches)

	req := httptest.NewRequest(http.MethodPost, CompanyControllerUrlQuery, nil)
	utils.SetHeader(req)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(10101), vjson.Parse(w.Body.String()).Get("code").Int())
}

func TestCompanyControllerQueryErrorBindBodyMock(t *testing.T) {
	g, c, patches := initCompanyController()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(c.BaseController, "BindBody", errors.New("bind body error"))
	defer vmock.Reset(patch)

	body, _ := json.Marshal(req.CompanyQueryReq{TraceIdReq: vreq.TraceIdReq{TraceId: "test"}, Id: 1})
	req := httptest.NewRequest(http.MethodPost, CompanyControllerUrlQuery, bytes.NewBuffer(body))
	utils.SetHeader(req)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, "", w.Body.String())
}
