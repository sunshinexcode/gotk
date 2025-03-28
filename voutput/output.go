package voutput

import (
	"net/http"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vvar"
)

var (
	E = Error
	O = Output
	S = Success
)

func Error(c *vapi.Context, err error, httpStatus ...int) *vapi.Context {
	if len(httpStatus) == 0 {
		httpStatus = append(httpStatus, http.StatusOK)
	}

	errCode, errMsg := verror.GetCode(err)
	vlog.Errorc(c, "Error", "api", c.FullPath(), "errCode", errCode, "errMsg", errMsg, "errStack", err, "httpStatus", httpStatus)

	c.Set("code", errCode)
	c.JSON(httpStatus[0], vapi.H{"code": errCode, "msg": errMsg, "data": vmap.M{}})

	return c
}

func Output(c *vapi.Context, err error, data any, httpStatus ...int) {
	if err != nil {
		Error(c, err, httpStatus...)
		return
	}

	Success(c, vcode.CodeSuccess.Message(), data)
}

func Success(c *vapi.Context, msg string, data any) {
	if vvar.IsNil(data) {
		data = vmap.M{}
	}

	c.Set("code", vcode.CodeSuccess.Code())
	c.JSON(http.StatusOK, vapi.H{"code": vcode.CodeSuccess.Code(), "msg": msg, "data": data})
}
