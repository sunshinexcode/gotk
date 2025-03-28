package vcontroller_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcontroller"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestBaseControllerBindBodyMock(t *testing.T) {
	baseController := vcontroller.NewBaseController(vcontroller.BaseControllerParam{})
	w := httptest.NewRecorder()
	c, _ := vapi.CreateTestContext(w)

	patch := vmock.ApplyMethodReturn(c, "ShouldBindBodyWith", errors.New("bind body error"))
	defer vmock.Reset(patch)

	err := baseController.BindBody(c, nil)

	vtest.NotNil(t, err)
	vtest.Equal(t, "bind body error", err.Error())
}
