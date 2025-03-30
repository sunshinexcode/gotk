package vmask

import (
	"fmt"
	"reflect"
	"strings"

	masker "github.com/ggwhite/go-masker/v2"

	"github.com/sunshinexcode/gotk/vlog"
)

// MaskStruct masks sensitive fields in a struct
// It is a wrapper function that provides a simpler interface for masking operations
// by handling error logging internally.
//
// val: The input struct or map to be masked
//
// Return: The masked struct with sensitive fields replaced by mask patterns
//
// Example:
//
//	type User struct {
//	    Password string `mask:"secret"`
//	}
//	user := &User{Password: "123456"}
//	masked := MaskStruct(user)
//	// Result: &User{Password: "12345*****"}
func MaskStruct(val any) any {
	s, err := MaskStructAndMap(val)
	if err != nil {
		vlog.Error("MaskStructAndMap", "err", err)
	}

	return s
}

// MaskStructAndMap masks sensitive fields in a struct or map
// maskKey specifies the keys to be masked, separated by |
func MaskStructAndMap(s any) (any, error) {
	sm, err := Mask.Struct(s)
	if err != nil {
		return sm, err
	}

	var selem reflect.Value
	st := reflect.TypeOf(sm)

	if st.Kind() == reflect.Ptr {
		selem = reflect.ValueOf(sm).Elem()
	} else {
		selem = reflect.ValueOf(sm)
	}

	tptr := reflect.New(selem.Type())
	tptr.Elem().Set(selem)

	for i := 0; i < selem.NumField(); i++ {
		field := tptr.Elem().Field(i)
		fieldType := selem.Type().Field(i)

		maskTag := fieldType.Tag.Get("mask")
		maskKeys := strings.Split(fieldType.Tag.Get("maskKey"), "|")

		if maskTag != "" {
			switch field.Kind() {
			case reflect.Struct:
				if masker.MaskerType(maskTag) == masker.MaskerTypeStruct {
					newVal, err := MaskStructAndMap(field.Interface())
					if err != nil {
						return nil, err
					}

					if reflect.TypeOf(newVal).Kind() == reflect.Ptr {
						field.Set(reflect.ValueOf(newVal).Elem())
					} else {
						field.Set(reflect.ValueOf(newVal))
					}
				}

			case reflect.Map:
				keySet := make(map[string]struct{})
				for _, k := range maskKeys {
					if k != "" {
						keySet[k] = struct{}{}
					}
				}

				tmpMap := reflect.MakeMap(field.Type())
				iter := field.MapRange()

				for iter.Next() {
					key := iter.Key()
					mapVal := iter.Value()
					keyStr := fmt.Sprintf("%v", key.Interface())

					if mapVal.Kind() == reflect.Interface {
						mapVal = mapVal.Elem()
					}

					var processedVal reflect.Value
					if _, ok := keySet[keyStr]; ok && mapVal.Kind() == reflect.String {
						masked := MaskSecret(mapVal.String(), MaskKeepLen, MaskPattern)
						processedVal = reflect.ValueOf(masked)
					} else {
						processedVal = mapVal
					}

					tmpMap.SetMapIndex(key, processedVal)
				}
				field.Set(tmpMap)
			}
		}
	}

	return tptr.Interface(), nil
}
