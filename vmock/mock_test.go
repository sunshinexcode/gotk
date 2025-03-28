package vmock_test

import (
	"reflect"
	"testing"

	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vredis"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtime"
)

type User struct {
}

func (u *User) eat() string {
	return "apple"
}

var (
	fn = func() string {
		return "0"
	}

	num = 1
)

func TestApplyFunc(t *testing.T) {
	patch := vmock.ApplyFunc(vtime.GetNowUtcUnix, func() int64 {
		return int64(1)
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, int64(1), vtime.GetNowUtcUnix())
}

func TestApplyFuncReturn(t *testing.T) {
	patch := vmock.ApplyFuncReturn(vtime.GetNowUtcUnix, int64(1))
	defer vmock.Reset(patch)

	vtest.Equal(t, int64(1), vtime.GetNowUtcUnix())
}

func TestApplyFuncSeq(t *testing.T) {
	patch := vmock.ApplyFuncSeq(vtime.GetNowUtcUnix, []vmock.OutputCell{{Values: vmock.Params{int64(1)}}, {Values: vmock.Params{int64(2)}}})
	defer vmock.Reset(patch)

	vtest.Equal(t, int64(1), vtime.GetNowUtcUnix())
	vtest.Equal(t, int64(2), vtime.GetNowUtcUnix())
}

func TestApplyFuncVar(t *testing.T) {
	patch := vmock.ApplyFuncVar(&fn, func() string {
		return "test"
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, "test", fn())
}

func TestApplyFuncVarReturn(t *testing.T) {
	patch := vmock.ApplyFuncVarReturn(&fn, "test")
	defer vmock.Reset(patch)

	vtest.Equal(t, "test", fn())
}

func TestApplyFuncVarSeq(t *testing.T) {
	patch := vmock.ApplyFuncVarSeq(&fn, []vmock.OutputCell{{Values: vmock.Params{"test1"}}, {Values: vmock.Params{"test2"}}})
	defer vmock.Reset(patch)

	vtest.Equal(t, "test1", fn())
	vtest.Equal(t, "test2", fn())
}

func TestApplyGlobalVar(t *testing.T) {
	patch := vmock.ApplyGlobalVar(&num, 2)
	defer vmock.Reset(patch)

	vtest.Equal(t, 2, num)
}

func TestApplyMethod(t *testing.T) {
	redis := vredis.Redis{}

	patch := vmock.ApplyMethod(reflect.TypeOf(&redis), "Ping", func(_ *vredis.Redis) (string, error) {
		return "pong", nil
	})
	defer vmock.Reset(patch)

	res, err := redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "pong", res)
}

func TestApplyMethodFunc(t *testing.T) {
	redis := vredis.Redis{}

	patch := vmock.ApplyMethodFunc(reflect.TypeOf(&redis), "Ping", func() (string, error) {
		return "pong", nil
	})
	defer vmock.Reset(patch)

	res, err := redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "pong", res)
}

func TestApplyMethodReturn(t *testing.T) {
	redis := vredis.Redis{}

	patch := vmock.ApplyMethodReturn(&redis, "Ping", "pong", nil)
	defer vmock.Reset(patch)

	res, err := redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "pong", res)
}

func TestApplyMethodSeq(t *testing.T) {
	redis := vredis.Redis{}

	patch := vmock.ApplyMethodSeq(&redis, "Ping", []vmock.OutputCell{{Values: vmock.Params{"pong1", nil}}, {Values: vmock.Params{"pong2", nil}}})
	defer vmock.Reset(patch)

	res, err := redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "pong1", res)

	res, err = redis.Ping()

	vtest.Nil(t, err)
	vtest.Equal(t, "pong2", res)
}

func TestApplyPrivateMethod(t *testing.T) {
	user := User{}

	patch := vmock.ApplyPrivateMethod(&user, "eat", func() string {
		return "pineapple"
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, "pineapple", user.eat())
}

func TestResetMock(t *testing.T) {
	patches := make([]*vmock.Patches, 0)
	patches = append(patches, vmock.ApplyFuncVarReturn(&fn, "test"))
	patches = append(patches, vmock.ApplyGlobalVar(&num, 2))

	vmock.ResetAll(patches)
}
