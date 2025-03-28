package bootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/bootstrap"
)

func TestNewHttpServerMock(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vbootstrap.NewHttpServer, &vapi.Engine{})
	defer vmock.Reset(patch)

	e := bootstrap.NewHttpServer(&vfx.LifecycleMock{}, &configs.Config{})

	vtest.Equal(t, "", e.BasePath())
}
