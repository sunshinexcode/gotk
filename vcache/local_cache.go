package vcache

import (
	"context"
	"time"

	"github.com/sunshinexcode/gotk/vvar"
)

var _ ILocalCache = (*LocalCache)(nil)

type ILocalCache interface {
	Get(ctx context.Context, key interface{}) (*vvar.Var, error)
	Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
}

type LocalCache struct {
}

func NewLocalCache() ILocalCache {
	return &LocalCache{}
}

func (r *LocalCache) Get(ctx context.Context, key interface{}) (*vvar.Var, error) {
	return Get(ctx, key)
}

func (r *LocalCache) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return Set(ctx, key, value, duration)
}
