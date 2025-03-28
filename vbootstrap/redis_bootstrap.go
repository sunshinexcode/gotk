package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vredis"
)

// NewRedis new redis
func NewRedis(config vconfig.IConfig) (redis *vredis.Redis, err error) {
	var options vmap.M
	redis = &vredis.Redis{}

	defer func() {
		if err != nil {
			vlog.Error("NewRedis", "err", err, "options", redis.Options)
		}
	}()

	if err = vmap.Decode(config.GetRedis(), &options); err != nil {
		return
	}
	if redis, err = vredis.New(options); err != nil {
		return
	}

	vlog.Infof("NewRedis, addrs:%+v, db:%d", redis.Options.Addrs, redis.Options.DB)
	return
}
