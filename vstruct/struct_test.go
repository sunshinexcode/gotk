package vstruct_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vstruct"
	"github.com/sunshinexcode/gotk/vtest"
)

type (
	User struct {
		Name   string
		Age    int32
		Salary int
	}

	Employee struct {
		Name   string
		Age    int32
		Salary int
	}
)

func TestCopy(t *testing.T) {
	var a vmap.M

	err := vstruct.Copy(&a, &vmap.M{"a": 1})

	vtest.Nil(t, err)
	vtest.Equal(t, vmap.M{"a": 1}, a)
}

func TestCopyError(t *testing.T) {
	var a int

	err := vstruct.Copy(a, &vmap.M{"a": 1})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, errors.Is(err, verror.ErrDataCopy))
	vtest.Equal(t, "10111|data copy error|0 &map[a:1]  \n-> copy destination must be non-nil and addressable", err.Error())
}

func TestCopyWithOption(t *testing.T) {
	var a vmap.M

	err := vstruct.CopyWithOption(&a, vmap.M{"a": 1, "b": ""}, vstruct.Option{IgnoreEmpty: true})

	vtest.Nil(t, err)
	vtest.Equal(t, vmap.M{"a": 1, "b": ""}, a)

	user := User{
		Name:   "",
		Age:    18,
		Salary: 2000,
	}
	employee := Employee{Name: "gotk"}

	err = vstruct.CopyWithOption(&employee, user, vstruct.Option{IgnoreEmpty: true})

	vtest.Nil(t, err)
	vtest.Equal(t, Employee{Name: "gotk", Age: 18, Salary: 2000}, employee)

	user = User{
		Name:   "",
		Age:    18,
		Salary: 2000,
	}
	employee = Employee{Name: "gotk"}

	err = vstruct.CopyWithOption(&employee, user, vstruct.Option{IgnoreEmpty: false})

	vtest.Nil(t, err)
	vtest.Equal(t, Employee{Name: "", Age: 18, Salary: 2000}, employee)
}

func TestCopyWithOptionError(t *testing.T) {
	var a int

	err := vstruct.CopyWithOption(a, &vmap.M{"a": 1}, vstruct.Option{})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, errors.Is(err, verror.ErrDataCopy))
	vtest.Equal(t, "10111|data copy error|0 &map[a:1] {IgnoreEmpty:false CaseSensitive:false DeepCopy:false Converters:[] FieldNameMapping:[]}  \n-> copy destination must be non-nil and addressable", err.Error())
}
