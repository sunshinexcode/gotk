package cache

import (
	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vfx"
)

var Module = vfx.Provide(
	vcache.NewLocalCache,
	vcache.NewRedisCache,
)
