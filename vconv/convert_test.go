package vconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestInt(t *testing.T) {
	vtest.Equal(t, 123, vconv.Int("123"))
	vtest.Equal(t, 123, vconv.Int(123))
	vtest.Equal(t, 123, vconv.Int(123.0))
	vtest.Equal(t, 123, vconv.Int(123.12))
}

func TestInt64(t *testing.T) {
	vtest.Equal(t, int64(123), vconv.Int64("123"))
	vtest.Equal(t, int64(123), vconv.Int64(123))
	vtest.Equal(t, int64(123), vconv.Int64(123.0))
	vtest.Equal(t, int64(123), vconv.Int64(123.12))
}

func TestMap(t *testing.T) {
	type User struct {
		Uid  int    `c:"uid"`
		Name string `c:"name"`
	}

	vtest.Equal(t, "john", vconv.Map(User{Uid: 1, Name: "john"})["name"])
	vtest.Equal(t, "john", vconv.Map(&User{Uid: 1, Name: "john"})["name"])

	type User2 struct {
		Uid  int
		Name string
	}

	vtest.Equal(t, nil, vconv.Map(User2{Uid: 1, Name: "john"})["name"])

	type User3 struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}

	vtest.Equal(t, "john", vconv.Map(User3{Uid: 1, Name: "john"})["name"])

	type User4 struct {
		Uid  int    `form:"uid"`
		Name string `form:"name"`
	}

	vtest.Equal(t, "john", vconv.Map(User4{Uid: 1, Name: "john"}, gconv.MapOption{Tags: []string{"form"}})["name"])

	type Base struct {
		Id   int    `c:"id"`
		Date string `c:"date"`
	}
	type User5 struct {
		UserBase Base   `c:"base"`
		Uid      int    `c:"uid"`
		Name     string `c:"name"`
	}

	vtest.Equal(t, 1, vconv.Map(User5{UserBase: Base{Id: 1}, Uid: 1, Name: "john"})["base"].(Base).Id)
}

func TestString(t *testing.T) {
	vtest.Equal(t, "123", vconv.String("123"))
	vtest.Equal(t, "123", vconv.String(123))
	vtest.Equal(t, "123", vconv.String(123.0))
	vtest.Equal(t, "123.12", vconv.String(123.12))
}

func TestStruct(t *testing.T) {
	type User struct {
		Uid  int
		Name string
	}
	params := vmap.M{
		"uid":  1,
		"name": "john",
	}
	var user *User
	err := vconv.Struct(params, &user)

	vtest.Nil(t, err)
	vtest.Equal(t, 1, user.Uid)
	vtest.Equal(t, "john", user.Name)

	type Ids struct {
		Id  int `json:"id"`
		Uid int `json:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `json:"create_time"`
	}
	type User2 struct {
		Base
		Passport string `json:"passport"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	data := vmap.M{
		"id":          1,
		"uid":         100,
		"passport":    "john",
		"password":    "123456",
		"nickname":    "John",
		"create_time": "2019",
	}
	user2 := new(User2)
	err = vconv.Struct(data, user2)

	vtest.Nil(t, err)
	vtest.Equal(t, 1, user2.Id)
	vtest.Equal(t, 100, user2.Uid)
	vtest.Equal(t, "John", user2.Nickname)
}
