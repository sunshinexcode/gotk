package vreflect_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vtest"
)

type Student struct {
	Name string
	age  int
}

func TestSetAttr(t *testing.T) {
	student := &Student{Name: "test", age: 10}
	err := vreflect.SetAttr(student, "Name", "test_modify")

	vtest.Nil(t, err)

	err = vreflect.SetAttr(student, "age", 20)

	vtest.NotNil(t, err)
	vtest.Equal(t, "cannot set attr value, attr:age", err.Error())
	vtest.Equal(t, "test_modify", student.Name)
	vtest.Equal(t, 10, student.age)

	err = vreflect.SetAttr(student, "Name", 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "error type, attr:Name, wrong:string, correct:int", err.Error())

	err = vreflect.SetAttr(student, "sex", 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:sex", err.Error())
}

func TestSetAttrs(t *testing.T) {
	student := &Student{Name: "test", age: 10}
	err := vreflect.SetAttrs(student, map[string]any{"Name": "test_modify", "age": 20})

	vtest.NotNil(t, err)
	vtest.Equal(t, "test_modify", student.Name)
	vtest.Equal(t, 10, student.age)
}
