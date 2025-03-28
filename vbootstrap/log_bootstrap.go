package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
)

// NewLog new log
func NewLog(config vconfig.IConfig) (err error) {
	var options vmap.M

	defer func() {
		if err != nil {
			vlog.Error("NewLog", "err", err, "options", vlog.GetLog().Options)
		}
	}()

	if err = vmap.Decode(config.GetLog(), &options); err != nil {
		return
	}
	if _, err = vlog.SetConfig(options); err != nil {
		return
	}

	vlog.Infof("NewLog, options:%+v", vlog.GetLog().Options)
	return
}
