package vredis_test

import (
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vredis"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestC(t *testing.T) {
	redis := &vredis.Redis{}

	vtest.Nil(t, redis.C())
}

func TestPing(t *testing.T) {
	redis, err := vredis.New(vmap.M{"Addrs": []string{"localhost:1"}})

	vtest.Nil(t, err)

	_, err = redis.Ping()

	vtest.NotNil(t, err)
}

func TestPingMock(t *testing.T) {
	var redis *vredis.Redis
	patch := vmock.ApplyMethod(reflect.TypeOf(redis), "Ping", func(_ *vredis.Redis) (string, error) {
		return "PONG", nil
	})
	defer vmock.Reset(patch)

	redis, err := vredis.New(nil)

	vtest.Nil(t, err)
	pong, err := redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "PONG", pong)
}

func TestSetConfig(t *testing.T) {
	redis := &vredis.Redis{Options: &redis.UniversalOptions{}}
	err := redis.SetConfig(vmap.M{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
}
