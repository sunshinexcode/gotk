package vpprof

import (
	_ "net/http/pprof"

	"github.com/sunshinexcode/gotk/vfx"
)

var Module = vfx.Options(
	vfx.Invoke(New),
)
