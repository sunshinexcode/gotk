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

// Ensure RedisCache implements IRedisCache interface
var _ IRedisCache = (*RedisCache)(nil)

// IRedisCache defines the interface for Redis cache operations
type IRedisCache interface {
	// Del deletes one or more keys from Redis
	// Returns the number of keys that were removed
	Del(ctx context.Context, keys ...string) (int64, error)

	// ExpireNX sets a key's time to live in seconds if the key exists
	// Returns true if the timeout was set, false if the key does not exist
	ExpireNX(ctx context.Context, key string, duration time.Duration) (bool, error)

	// Get retrieves the value of a key
	// Returns the value of the key, or empty string if the key does not exist
	Get(ctx context.Context, key string) (string, error)

	// HDel deletes one or more hash fields
	// Returns the number of fields that were removed
	HDel(ctx context.Context, key string, fields ...string) (int64, error)

	// HGet retrieves the value of a hash field
	// Returns the value associated with the field, or empty string if the field does not exist
	HGet(ctx context.Context, key string, field string) (string, error)

	// HGetAll retrieves all fields and values of a hash
	// Returns a map of fields and their values
	HGetAll(ctx context.Context, key string) (map[string]string, error)

	// HSet sets the string value of a hash field
	// Returns the number of fields that were added
	HSet(ctx context.Context, key string, values ...any) (int64, error)

	// Lock acquires a distributed lock
	// Returns true if the lock was acquired, false if the lock is already held
	Lock(ctx context.Context, lockName string, duration time.Duration) (res bool, err error)

	// SAdd adds one or more members to a set
	// Returns the number of elements that were added to the set
	SAdd(ctx context.Context, key string, members ...any) (int64, error)

	// SPop removes and returns one random member from a set
	// Returns the removed member, or empty string if the set is empty
	SPop(ctx context.Context, key string) (string, error)

	// Set sets the string value of a key
	// Returns "OK" if the operation was successful
	Set(ctx context.Context, key string, value any, duration time.Duration) (string, error)

	// SetNX sets the value of a key, only if the key does not exist
	// Returns true if the key was set, false if the key already exists
	SetNX(ctx context.Context, key string, value any, duration time.Duration) (bool, error)

	// Unlock releases a distributed lock
	// Returns the number of keys that were removed
	Unlock(ctx context.Context, lockName string) (int64, error)
}

// RedisCacheParam contains the dependencies required for RedisCache
type RedisCacheParam struct {
	vfx.In
	Redis *vredis.Redis
}

// RedisCache implements the IRedisCache interface
type RedisCache struct {
	Redis *vredis.Redis
}

// NewRedisCache creates a new RedisCache instance
func NewRedisCache(p RedisCacheParam) IRedisCache {
	return &RedisCache{Redis: p.Redis}
}

// Del deletes one or more keys from Redis
func (r *RedisCache) Del(ctx context.Context, keys ...string) (int64, error) {
	res, err := r.Redis.C().Del(ctx, keys...).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, keys, res)
}

// ExpireNX sets a key's time to live in seconds if the key exists
func (r *RedisCache) ExpireNX(ctx context.Context, key string, duration time.Duration) (bool, error) {
	res, err := r.Redis.C().ExpireNX(ctx, key, duration).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// Get retrieves the value of a key
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	res, err := r.Redis.C().Get(ctx, key).Result()
	if err == redis.Nil {
		err = nil
	}
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// HDel deletes one or more hash fields
func (r *RedisCache) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	res, err := r.Redis.C().HDel(ctx, key, fields...).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// HGet retrieves the value of a hash field
func (r *RedisCache) HGet(ctx context.Context, key string, field string) (string, error) {
	res, err := r.Redis.C().HGet(ctx, key, field).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// HGetAll retrieves all fields and values of a hash
func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	res, err := r.Redis.C().HGetAll(ctx, key).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// HSet sets the string value of a hash field
func (r *RedisCache) HSet(ctx context.Context, key string, values ...any) (int64, error) {
	res, err := r.Redis.C().HSet(ctx, key, values...).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// Lock acquires a distributed lock
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

// SAdd adds one or more members to a set
func (r *RedisCache) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
	res, err := r.Redis.C().SAdd(ctx, key, members...).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// SPop removes and returns one random member from a set
func (r *RedisCache) SPop(ctx context.Context, key string) (string, error) {
	res, err := r.Redis.C().SPop(ctx, key).Result()
	if err == redis.Nil {
		err = nil
	}
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// Set sets the string value of a key
func (r *RedisCache) Set(ctx context.Context, key string, value any, duration time.Duration) (string, error) {
	res, err := r.Redis.C().Set(ctx, key, value, duration).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// SetNX sets the value of a key, only if the key does not exist
func (r *RedisCache) SetNX(ctx context.Context, key string, value any, duration time.Duration) (bool, error) {
	res, err := r.Redis.C().SetNX(ctx, key, value, duration).Result()
	return res, verror.Wrap(err, vcode.CodeErrRedisOperation, key, res)
}

// Unlock releases a distributed lock
func (r *RedisCache) Unlock(ctx context.Context, lockName string) (int64, error) {
	return r.Del(ctx, lockName)
}
