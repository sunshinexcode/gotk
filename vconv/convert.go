package vconv

import "github.com/gogf/gf/v2/util/gconv"

type (
	MapOption = gconv.MapOption
)

// Int converts `any` to int.
func Int(any any) int {
	return gconv.Int(any)
}

// Int64 converts `any` to int64.
func Int64(any any) int64 {
	return gconv.Int64(any)
}

// Map converts any variable `value` to map[string]any. If the parameter `value` is not a
// map/struct/*struct type, then the conversion will fail and returns nil.
//
// If `value` is a struct/*struct object, the second parameter `priorityTagAndFieldName` specifies the most priority
// priorityTagAndFieldName that will be detected, otherwise it detects the priorityTagAndFieldName in order of:
// gconv, json, field name.
func Map(value any, tags ...MapOption) map[string]any {
	return gconv.Map(value, tags...)
}

// String converts `any` to string.
// It's most commonly used converting function.
func String(any any) string {
	return gconv.String(any)
}

// Struct maps the params key-value pairs to the corresponding struct object's attributes.
// The third parameter `mapping` is unnecessary, indicating the mapping rules between the
// custom key name and the attribute name(case-sensitive).
//
// Note:
//  1. The `params` can be any type of map/struct, usually a map.
//  2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object
//     or struct pointer.
//  3. Only the public attributes of struct object can be mapped.
//  4. If `params` is a map, the key of the map `params` can be lowercase.
//     It will automatically convert the first letter of the key to uppercase
//     in mapping procedure to do the matching.
//     It ignores the map key, if it does not match.
func Struct(params any, pointer any, mapping ...map[string]string) (err error) {
	return gconv.Struct(params, pointer, mapping...)
}
