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

// TestContains tests the Contains function
func TestContains(t *testing.T) {
	// Test non-existent key
	exists, err := vcache.Contains(context.TODO(), "key_contains")

	vtest.Nil(t, err)
	vtest.Equal(t, false, exists)

	// Test existing key
	err = vcache.Set(context.TODO(), "key_contains", "val_contains", 0)

	vtest.Nil(t, err)

	exists, err = vcache.Contains(context.TODO(), "key_contains")

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)

	// Test with different key types
	err = vcache.Set(context.TODO(), 123, "val_contains_int", 0)

	vtest.Nil(t, err)

	exists, err = vcache.Contains(context.TODO(), 123)

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)

	// Test with struct key
	type testKey struct {
		ID   int
		Name string
	}
	key := testKey{ID: 1, Name: "test"}
	err = vcache.Set(context.TODO(), key, "val_contains_struct", 0)

	vtest.Nil(t, err)

	exists, err = vcache.Contains(context.TODO(), key)

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)
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

func TestRemove(t *testing.T) {
	// Test removing non-existent key
	val, err := vcache.Remove(context.TODO(), "key_remove")

	vtest.Nil(t, err)
	vtest.Equal(t, "", val.String())

	// Test removing existing key
	err = vcache.Set(context.TODO(), "key_remove", "val_remove", 0)

	vtest.Nil(t, err)

	exists, err := vcache.Contains(context.TODO(), "key_remove")

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)

	val, err = vcache.Remove(context.TODO(), "key_remove")

	vtest.Nil(t, err)
	vtest.Equal(t, "val_remove", val.String())

	exists, err = vcache.Contains(context.TODO(), "key_remove")

	vtest.Nil(t, err)
	vtest.Equal(t, false, exists)

	// Test removing with different key types
	err = vcache.Set(context.TODO(), 123, "val_remove_int", 0)

	vtest.Nil(t, err)

	exists, err = vcache.Contains(context.TODO(), 123)

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)

	val, err = vcache.Remove(context.TODO(), 123)

	vtest.Nil(t, err)
	vtest.Equal(t, "val_remove_int", val.String())

	exists, err = vcache.Contains(context.TODO(), 123)

	vtest.Nil(t, err)
	vtest.Equal(t, false, exists)

	// Test removing struct key
	type testKey struct {
		ID   int
		Name string
	}
	key := testKey{ID: 1, Name: "test"}
	err = vcache.Set(context.TODO(), key, "val_remove_struct", 0)

	vtest.Nil(t, err)

	exists, err = vcache.Contains(context.TODO(), key)

	vtest.Nil(t, err)
	vtest.Equal(t, true, exists)

	val, err = vcache.Remove(context.TODO(), key)

	vtest.Nil(t, err)
	vtest.Equal(t, "val_remove_struct", val.String())

	exists, err = vcache.Contains(context.TODO(), key)

	vtest.Nil(t, err)
	vtest.Equal(t, false, exists)
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
