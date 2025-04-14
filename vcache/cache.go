package vcache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/sunshinexcode/gotk/vvar"
)

const (
	DataEmpty = ""
)

var (
	Contains = gcache.Contains
	Set      = gcache.Set
)

// CheckDataEmpty checks if the data is empty
func CheckDataEmpty(data string) bool {
	return data == DataEmpty
}

// Get retrieves and returns the associated value of given `key`.
// It returns nil if it does not exist, or its value is nil, or it's expired.
// If you would like to check if the `key` exists in the cache, it's better using function Contains.
func Get(ctx context.Context, key any) (*vvar.Var, error) {
	return gcache.Get(ctx, key)
}

// GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and
// returns `value` if `key` does not exist in the cache. The key-value pair expires
// after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
// if `value` is a function and the function result is nil.
func GetOrSet(ctx context.Context, key any, value any, duration time.Duration) (*vvar.Var, error) {
	return gcache.GetOrSet(ctx, key, value, duration)
}

// Remove deletes one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the last deleted item.
func Remove(ctx context.Context, keys ...any) (oldValue *vvar.Var, err error) {
	return gcache.Remove(ctx, keys...)
}

// Update updates the value of `key` without changing its expiration and returns the old value.
// The returned value `exist` is false if the `key` does not exist in the cache.
//
// It deletes the `key` if given `value` is nil.
// It does nothing if `key` does not exist in the cache.
func Update(ctx context.Context, key any, value any) (oldValue *vvar.Var, exist bool, err error) {
	return gcache.Update(ctx, key, value)
}
