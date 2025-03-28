package voutput_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestError(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())
	voutput.E(c, verror.ErrParamInvalid)
	voutput.Error(c, verror.ErrParamInvalid)

	vtest.Equal(t, 200, c.Writer.Status())

	c, _ = vapi.CreateTestContext(httptest.NewRecorder())
	voutput.E(c, verror.ErrParamInvalid, http.StatusBadGateway)
	voutput.Error(c, verror.ErrParamInvalid, http.StatusBadGateway)

	vtest.Equal(t, http.StatusBadGateway, c.Writer.Status())

	c, _ = vapi.CreateTestContext(httptest.NewRecorder())
	voutput.Error(c, verror.ErrParamInvalid, http.StatusBadGateway, http.StatusForbidden)

	vtest.Equal(t, http.StatusBadGateway, c.Writer.Status())
}

func TestOutput(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())
	voutput.O(c, nil, "test")
	voutput.Output(c, nil, "test")

	vtest.Equal(t, 200, c.Writer.Status())
}

func TestOutputError(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())
	voutput.O(c, verror.ErrHttpRequest, "test")

	vtest.Equal(t, 200, c.Writer.Status())
}

func TestSuccess(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())
	voutput.S(c, "", "test")
	voutput.Success(c, "", "test")

	vtest.Equal(t, 200, c.Writer.Status())
}

func TestSuccessNil(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())
	voutput.S(c, "", nil)

	vtest.Equal(t, 200, c.Writer.Status())
}
