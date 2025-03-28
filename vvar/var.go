package vvar

import "reflect"

func IsNil(value any) bool {
	if value == nil {
		return true
	}

	r := reflect.ValueOf(value)
	if r.Kind() == reflect.Ptr {
		return r.IsNil()
	}

	return false
}
