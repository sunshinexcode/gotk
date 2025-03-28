package vcache_test

import (
	"context"
	"testing"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCheckDataEmpty(t *testing.T) {
	vtest.Equal(t, true, vcache.CheckDataEmpty(""))
	vtest.Equal(t, false, vcache.CheckDataEmpty("test"))
}

func TestGet(t *testing.T) {
	data, err := vcache.Get(context.TODO(), "key_get")

	vtest.Nil(t, err)
	vtest.Equal(t, "", data.String())
}

func TestGetOrSet(t *testing.T) {
	data, err := vcache.GetOrSet(context.TODO(), "key_get_or_set", "val_get_or_set", 0)

	vtest.Nil(t, err)
	vtest.Equal(t, "val_get_or_set", data.String())

	data, err = vcache.GetOrSet(context.TODO(), "key_get_or_set", "val_get_or_set_2", 0)

	vtest.Nil(t, err)
	vtest.Equal(t, "val_get_or_set", data.String())
}

func TestSet(t *testing.T) {
	err := vcache.Set(context.TODO(), "key_set", "val_set", 0)

	vtest.Nil(t, err)

	data, err := vcache.Get(context.TODO(), "key_set")

	vtest.Nil(t, err)
	vtest.Equal(t, "val_set", data.String())

	err = vcache.Set(context.TODO(), "key_set", "val_set_2", 0)

	vtest.Nil(t, err)

	data, err = vcache.Get(context.TODO(), "key_set")

	vtest.Nil(t, err)
	vtest.Equal(t, "val_set_2", data.String())
}

func TestUpdate(t *testing.T) {
	data, existed, err := vcache.Update(context.TODO(), "key_update", "val_update")

	vtest.Nil(t, err)
	vtest.Equal(t, false, existed)
	vtest.Equal(t, "", data.String())

	err = vcache.Set(context.TODO(), "key_update", "val_update", 0)

	vtest.Nil(t, err)

	data, err = vcache.Get(context.TODO(), "key_update")

	vtest.Nil(t, err)
	vtest.Equal(t, "val_update", data.String())

	data, existed, err = vcache.Update(context.TODO(), "key_update", "val_update_changed")

	vtest.Nil(t, err)
	vtest.Equal(t, true, existed)
	vtest.Equal(t, "val_update", data.String())

	data, err = vcache.Get(context.TODO(), "key_update")

	vtest.Nil(t, err)
	vtest.Equal(t, "val_update_changed", data.String())
}
