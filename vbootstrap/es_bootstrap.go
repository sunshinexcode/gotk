package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/ves"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
)

// NewEs new es
func NewEs(config vconfig.IConfig) (es *ves.Es, err error) {
	var options vmap.M
	es = &ves.Es{}

	defer func() {
		if err != nil {
			vlog.Error("NewEs", "err", err, "options", es.Options)
		}
	}()

	if err = vmap.Decode(config.GetEs(), &options); err != nil {
		return
	}
	if es, err = ves.New(options); err != nil {
		return
	}

	vlog.Infof("NewEs, options:%+v", es.Options)
	return
}
