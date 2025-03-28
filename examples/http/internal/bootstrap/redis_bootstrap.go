package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vredis"

	"app/configs"
)

// NewRedis new redis
func NewRedis(config *configs.Config) (redis *vredis.Redis, err error) {
	return vbootstrap.NewRedis(config)
}
