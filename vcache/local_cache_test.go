package vcache_test

import (
	"context"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vtest"
)

type UserQueryResp struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func TestLocalCacheGet(t *testing.T) {
	LocalCache := vcache.NewLocalCache()

	data, err := LocalCache.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "", data.String())
}

func TestLocalCacheSet(t *testing.T) {
	LocalCache := vcache.NewLocalCache()

	err := LocalCache.Set(context.TODO(), "test", "test", 10*time.Second)

	vtest.Nil(t, err)

	data, err := LocalCache.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "test", data.String())
}

func TestLocalCacheSetMap(t *testing.T) {
	LocalCache := vcache.NewLocalCache()

	err := LocalCache.Set(context.TODO(), "test", map[string]any{"name": "test", "age": 18}, 10*time.Second)

	vtest.Nil(t, err)

	data, err := LocalCache.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "{\"age\":18,\"name\":\"test\"}", data.String())
	vtest.Equal(t, false, data.IsEmpty())
	vtest.Equal(t, false, data.IsNil())
	vtest.Equal(t, "test", data.Map()["name"])
}

func TestLocalCacheSetStruct(t *testing.T) {
	LocalCache := vcache.NewLocalCache()

	err := LocalCache.Set(context.TODO(), "test", &UserQueryResp{Id: 1}, 10*time.Second)

	vtest.Nil(t, err)

	data, err := LocalCache.Get(context.TODO(), "test")

	vtest.Nil(t, err)
	vtest.Equal(t, "{\"id\":1}", data.String())
	vtest.Equal(t, false, data.IsEmpty())
	vtest.Equal(t, false, data.IsNil())

	dataS := &UserQueryResp{}

	vtest.Nil(t, data.Struct(dataS))
	vtest.Equal(t, int64(1), dataS.Id)
}
