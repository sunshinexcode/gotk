package vmap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestType(t *testing.T) {
	vtest.Equal(t, map[string]any{"name": "test"}, vmap.M{"name": "test"})
}

func TestDecode(t *testing.T) {
	// Struct -> Map
	var m vmap.M
	type User struct {
		Uid  int
		Name string
	}
	err := vmap.Decode(User{Name: "test"}, &m)

	vtest.Nil(t, err)
	vtest.Equal(t, map[string]any{"Name": "test", "Uid": 0}, m)
	vtest.Equal(t, "test", m["Name"])
	vtest.Equal(t, nil, m["name"])

	type User2 struct {
		Uid  int
		Name string `json:"name"`
	}
	err = vmap.Decode(User2{Name: "test"}, &m)

	vtest.Nil(t, err)
	vtest.Equal(t, map[string]any{"Name": "test", "Uid": 0}, m)
	vtest.Equal(t, "test", m["Name"])
	vtest.Equal(t, nil, m["name"])

	type User3 struct {
		Uid  int
		Name string `mapstructure:"name"`
	}
	err = vmap.Decode(User3{Name: "test"}, &m)

	vtest.Nil(t, err)
	vtest.Equal(t, map[string]any{"Name": "test", "Uid": 0, "name": "test"}, m)
	vtest.Equal(t, "test", m["Name"])
	vtest.Equal(t, "test", m["name"])

	// Struct -> Struct
	var user User
	err = vmap.Decode(User3{Name: "test", Uid: 1}, &user)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", user.Name)
	vtest.Equal(t, 1, user.Uid)

	// Map -> Struct
	var user2 User
	err = vmap.Decode(vmap.M{"name": "test", "uid": 2}, &user2)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", user2.Name)
	vtest.Equal(t, 2, user2.Uid)
}

func TestGetKeys(t *testing.T) {
	vtest.Equal(t, []string(nil), vmap.GetKeys(vmap.M{}))
	vtest.Equal(t, []string{"a"}, vmap.GetKeys(vmap.M{"a": 1}))
}

func TestMerge(t *testing.T) {
	mapData := vmap.Merge(map[string]any{"name": "t1", "age": 10}, map[string]any{})
	mapDataJson, _ := vjson.Encode(mapData)

	vtest.Equal(t, `{"age":10,"name":"t1"}`, mapDataJson)

	mapData = vmap.Merge(map[string]any{"name": "t1", "age": 10}, map[string]any{"name": "t2"})
	mapDataJson, _ = vjson.Encode(mapData)

	vtest.Equal(t, `{"age":10,"name":"t2"}`, mapDataJson)

	mapData = vmap.Merge(map[string]any{"name": "t1", "age": 10}, map[string]any{"name": "t2", "sex": "man"})
	mapDataJson, _ = vjson.Encode(mapData)

	vtest.Equal(t, `{"age":10,"name":"t2","sex":"man"}`, mapDataJson)

	mapData = vmap.Merge(map[string]any{"name": "t1", "age": 10}, map[string]any{"name": "t2", "sex": "man"}, map[string]any{"name": "t3", "sex": "woman"})
	mapDataJson, _ = vjson.Encode(mapData)

	vtest.Equal(t, `{"age":10,"name":"t3","sex":"woman"}`, mapDataJson)
}

func TestNew(t *testing.T) {
	m := vmap.New()
	m.Set("key1", "val1")
	m.Set("key2", "val2")

	vtest.Equal(t, "val1", m.Get("key1"))
	vtest.Equal(t, 2, m.Size())
	vtest.Equal(t, true, m.Contains("key1"))
	vtest.Equal(t, false, m.Contains("key3"))
	vtest.Equal(t, false, m.IsEmpty())
}

func TestSortKey(t *testing.T) {
	vtest.Equal(t, []string{"a", "b", "c"}, vmap.SortKey(vmap.M{"a": 1, "c": "2", "b": 3}))
	vtest.Equal(t, []string(nil), vmap.SortKey(vmap.M{}))
	vtest.Equal(t, []string{"a"}, vmap.SortKey(vmap.M{"a": 1}))
}
