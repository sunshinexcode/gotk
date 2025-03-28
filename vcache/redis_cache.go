package vcache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vredis"
)

var _ IRedisCache = (*RedisCache)(nil)

type IRedisCache interface {
	Del(ctx context.Context, keys ...string) (int64, error)
	ExpireNX(ctx context.Context, key string, duration time.Duration) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	Lock(ctx context.Context, lockName string, duration time.Duration) (res bool, err error)
	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	SPop(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) (string, error)
	SetNX(ctx context.Context, key string, value interface{}, duration time.Duration) (bool, error)
	Unlock(ctx context.Context, lockName string) (int64, error)
}

type RedisCacheParam struct {
	vfx.In

	Redis *vredis.Redis
}

type RedisCache struct {
	Redis *vredis.Redis
}

func NewRedisCache(p RedisCacheParam) IRedisCache {
	return &RedisCache{Redis: p.Redis}
}

func (r *RedisCache) Del(ctx context.Context, keys ...string) (int64, error) {
	res, err := r.Redis.C().Del(ctx, keys...).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, keys, res)
}

func (r *RedisCache) ExpireNX(ctx context.Context, key string, duration time.Duration) (bool, error) {
	res, err := r.Redis.C().ExpireNX(ctx, key, duration).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	res, err := r.Redis.C().Get(ctx, key).Result()

	if err == redis.Nil {
		err = nil
	}

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	res, err := r.Redis.C().HDel(ctx, key, fields...).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) HGet(ctx context.Context, key string, field string) (string, error) {
	res, err := r.Redis.C().HGet(ctx, key, field).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	res, err := r.Redis.C().HGetAll(ctx, key).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	res, err := r.Redis.C().HSet(ctx, key, values...).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) Lock(ctx context.Context, lockName string, duration time.Duration) (res bool, err error) {
	res, err = r.SetNX(ctx, lockName, 1, duration)

	if err != nil {
		return
	}

	if !res {
		return false, verror.ErrRedisAcquireLockFailed
	}

	return
}

func (r *RedisCache) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	res, err := r.Redis.C().SAdd(ctx, key, members...).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) SPop(ctx context.Context, key string) (string, error) {
	res, err := r.Redis.C().SPop(ctx, key).Result()

	if err == redis.Nil {
		err = nil
	}

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (string, error) {
	res, err := r.Redis.C().Set(ctx, key, value, duration).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) SetNX(ctx context.Context, key string, value interface{}, duration time.Duration) (bool, error) {
	res, err := r.Redis.C().SetNX(ctx, key, value, duration).Result()

	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

func (r *RedisCache) Unlock(ctx context.Context, lockName string) (int64, error) {
	return r.Del(ctx, lockName)
}
