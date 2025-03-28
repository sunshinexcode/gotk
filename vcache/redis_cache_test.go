package vcache_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vredis"
	"github.com/sunshinexcode/gotk/vtest"
)

func initRedisCache() (r vcache.IRedisCache, patches []*vmock.Patches) {
	r = vcache.NewRedisCache(vcache.RedisCacheParam{Redis: &vredis.Redis{Client: &redis.Client{}}})

	patchDel := vmock.ApplyMethodReturn(&redis.Client{}, "Del", &redis.IntCmd{})
	patches = append(patches, patchDel)

	patchExpireNX := vmock.ApplyMethodReturn(&redis.Client{}, "ExpireNX", &redis.BoolCmd{})
	patches = append(patches, patchExpireNX)

	patchGet := vmock.ApplyMethodReturn(&redis.Client{}, "Get", &redis.StringCmd{})
	patches = append(patches, patchGet)

	patchHDel := vmock.ApplyMethodReturn(&redis.Client{}, "HDel", &redis.IntCmd{})
	patches = append(patches, patchHDel)

	patchHGet := vmock.ApplyMethodReturn(&redis.Client{}, "HGet", &redis.StringCmd{})
	patches = append(patches, patchHGet)

	patchHGetAll := vmock.ApplyMethodReturn(&redis.Client{}, "HGetAll", &redis.MapStringStringCmd{})
	patches = append(patches, patchHGetAll)

	patchHSet := vmock.ApplyMethodReturn(&redis.Client{}, "HSet", &redis.IntCmd{})
	patches = append(patches, patchHSet)

	patchSAdd := vmock.ApplyMethodReturn(&redis.Client{}, "SAdd", &redis.IntCmd{})
	patches = append(patches, patchSAdd)

	patchSPop := vmock.ApplyMethodReturn(&redis.Client{}, "SPop", &redis.StringCmd{})
	patches = append(patches, patchSPop)

	patchSet := vmock.ApplyMethodReturn(&redis.Client{}, "Set", &redis.StatusCmd{})
	patches = append(patches, patchSet)

	patchSetNX := vmock.ApplyMethodReturn(&redis.Client{}, "SetNX", &redis.BoolCmd{})
	patches = append(patches, patchSetNX)

	return
}

func TestRedisCacheDelMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), nil)
	defer vmock.Reset(patch)

	res, err := r.Del(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), res)
}

func TestRedisCacheGetMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "test", nil)
	defer vmock.Reset(patch)

	data, err := r.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "test", data)
}

func TestRedisCacheGetErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "test", errors.New("redis error"))
	defer vmock.Reset(patch)

	data, err := r.Get(context.TODO(), "test")

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test test  \n-> redis error", err.Error())
	vtest.Equal(t, "test", data)
}

func TestRedisCacheGetErrorNilMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "", redis.Nil)
	defer vmock.Reset(patch)

	data, err := r.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, vcache.DataEmpty, data)
}

func TestRedisCacheHDelMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), nil)
	defer vmock.Reset(patch)

	data, err := r.HDel(context.TODO(), "test", "name")

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), data)
}

func TestRedisCacheHDelErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), errors.New("redis error"))
	defer vmock.Reset(patch)

	data, err := r.HDel(context.TODO(), "test", "name")

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test 1  \n-> redis error", err.Error())
	vtest.Equal(t, int64(1), data)
}

func TestRedisCacheExpireNXMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.BoolCmd{}, "Result", true, nil)
	defer vmock.Reset(patch)

	res, err := r.ExpireNX(context.TODO(), "test", 5*time.Minute)

	vtest.Nil(t, err)
	vtest.Equal(t, true, res)
}

func TestRedisCacheExpireNXErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.BoolCmd{}, "Result", true, errors.New("redis error"))
	defer vmock.Reset(patch)

	res, err := r.ExpireNX(context.TODO(), "test", 5*time.Minute)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test true  \n-> redis error", err.Error())
	vtest.Equal(t, true, res)
}

func TestRedisCacheHGetMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "test", nil)
	defer vmock.Reset(patch)

	data, err := r.HGet(context.TODO(), "test", "name")

	vtest.Nil(t, err)
	vtest.Equal(t, "test", data)
}

func TestRedisCacheHGetErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "test", errors.New("redis error"))
	defer vmock.Reset(patch)

	data, err := r.HGet(context.TODO(), "test", "name")

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test test  \n-> redis error", err.Error())
	vtest.Equal(t, "test", data)
}

func TestRedisCacheHGetAllMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.MapStringStringCmd{}, "Result", map[string]string{"name": "test", "age": "18"}, nil)
	defer vmock.Reset(patch)

	data, err := r.HGetAll(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, map[string]string{"age": "18", "name": "test"}, data)
}

func TestRedisCacheHGetAllErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.MapStringStringCmd{}, "Result", map[string]string{"name": "test", "age": "18"}, errors.New("redis error"))
	defer vmock.Reset(patch)

	data, err := r.HGetAll(context.TODO(), "test")

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test map[age:18 name:test]  \n-> redis error", err.Error())
	vtest.Equal(t, map[string]string{"age": "18", "name": "test"}, data)
}

func TestRedisCacheHSetMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), nil)
	defer vmock.Reset(patch)

	data, err := r.HSet(context.TODO(), "test", "name", "test")

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), data)
}

func TestRedisCacheHSetErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), errors.New("redis error"))
	defer vmock.Reset(patch)

	data, err := r.HSet(context.TODO(), "test", "name", "test")

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test 1  \n-> redis error", err.Error())
	vtest.Equal(t, int64(1), data)
}

func TestRedisCacheLockMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&vcache.RedisCache{}, "SetNX", true, nil)
	defer vmock.Reset(patch)

	res, err := r.Lock(context.TODO(), "test", 10*time.Second)

	vtest.Nil(t, err)
	vtest.Equal(t, true, res)
}

func TestRedisCacheLockErrorRedisOperationMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&vcache.RedisCache{}, "SetNX", false, errors.New("test"))
	defer vmock.Reset(patch)

	res, err := r.Lock(context.TODO(), "test", 10*time.Second)

	vtest.NotNil(t, err)
	vtest.Equal(t, "test", err.Error())
	vtest.Equal(t, false, res)
}

func TestRedisCacheLockErrorRedisAcquireLockFailedMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&vcache.RedisCache{}, "SetNX", false, nil)
	defer vmock.Reset(patch)

	res, err := r.Lock(context.TODO(), "test", 10*time.Second)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10141|redis acquire lock failed|<nil>", err.Error())
	vtest.Equal(t, false, res)
}

func TestRedisCacheSAddMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.IntCmd{}, "Result", int64(1), nil)
	defer vmock.Reset(patch)

	res, err := r.SAdd(context.TODO(), "test", 1)

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), res)
}

func TestRedisCacheSPopMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "1", nil)
	defer vmock.Reset(patch)

	res, err := r.SPop(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "1", res)
}

func TestRedisCacheSPopErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StringCmd{}, "Result", "1", redis.Nil)
	defer vmock.Reset(patch)

	res, err := r.SPop(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "1", res)
}

func TestRedisCacheSetMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StatusCmd{}, "Result", "test", nil)
	defer vmock.Reset(patch)

	res, err := r.Set(context.TODO(), "test", "test", 2*time.Second)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", res)
}

func TestRedisCacheSetErrorMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.StatusCmd{}, "Result", "test", errors.New("redis error"))
	defer vmock.Reset(patch)

	res, err := r.Set(context.TODO(), "test", "test", 2*time.Second)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|test test  \n-> redis error", err.Error())
	vtest.Equal(t, "test", res)
}

func TestRedisCacheSetNXMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&redis.BoolCmd{}, "Result", true, nil)
	defer vmock.Reset(patch)

	res, err := r.SetNX(context.TODO(), "test", 1, 2*time.Second)

	vtest.Nil(t, err)
	vtest.Equal(t, true, res)
}

func TestRedisCacheUnlockMock(t *testing.T) {
	r, patches := initRedisCache()
	defer vmock.ResetAll(patches)

	patch := vmock.ApplyMethodReturn(&vcache.RedisCache{}, "Del", int64(1), nil)
	defer vmock.Reset(patch)

	res, err := r.Unlock(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), res)
}
