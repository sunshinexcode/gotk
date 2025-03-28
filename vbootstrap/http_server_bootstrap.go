package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vhttp"
)

func NewHttpServer(lifecycle vfx.Lifecycle, config vconfig.IConfig) *vapi.Engine {
	vapi.SetMode(config.GetHttpServer().Model)
	s := &vhttp.Server{Addr: config.GetHttpServer().Address, G: vapi.Default(), StopHandler: func() {}}

	return vhttp.NewServer(lifecycle, s)
}
