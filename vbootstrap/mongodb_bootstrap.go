package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmongodb"
)

// NewMongodb new mongodb
func NewMongodb(config vconfig.IConfig) (mongodb *vmongodb.Mongodb, err error) {
	var options vmap.M
	mongodb = &vmongodb.Mongodb{}

	defer func() {
		if err != nil {
			vlog.Error("NewMongodb", "err", err, "options", mongodb.Options)
		}
	}()

	if err = vmap.Decode(config.GetMongodb(), &options); err != nil {
		return
	}
	if mongodb, err = vmongodb.New(options); err != nil {
		return
	}

	vlog.Infof("NewMongodb, err:%v, options:%+v", err, mongodb.Options)
	return
}
