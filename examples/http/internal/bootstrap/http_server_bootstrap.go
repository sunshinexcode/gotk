package bootstrap

import (
	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vfx"

	"app/configs"
)

func NewHttpServer(lifecycle vfx.Lifecycle, config *configs.Config) *vapi.Engine {
	return vbootstrap.NewHttpServer(lifecycle, config)
}
