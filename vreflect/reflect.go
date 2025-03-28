package vreflect

import (
	"fmt"
	"reflect"
)

// SetAttr assigns value to the object property
func SetAttr(obj any, name string, value any) (err error) {
	structAttrValue := reflect.ValueOf(obj).Elem().FieldByName(name)
	if !structAttrValue.IsValid() {
		return fmt.Errorf("no attr, attr:%s", name)
	}
	if !structAttrValue.CanSet() {
		return fmt.Errorf("cannot set attr value, attr:%s", name)
	}

	val := reflect.ValueOf(value)
	if structAttrValue.Type() != val.Type() {
		return fmt.Errorf("error type, attr:%s, wrong:%s, correct:%s", name, structAttrValue.Type(), val.Type())
	}

	structAttrValue.Set(val)
	return
}

// SetAttrs assigns value to the object properties
func SetAttrs(obj any, options map[string]any) (err error) {
	for k, v := range options {
		errAttr := SetAttr(obj, k, v)
		if errAttr != nil {
			err = errAttr
		}
	}

	return
}
