package vmap

import (
	"sort"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/mitchellh/mapstructure"
)

type (
	M   = map[string]any
	Map = map[string]any
)

// Decode takes an input structure and uses reflection to translate it to
// the output structure. output must be a pointer to a map or struct.
// Struct -> Map
// Struct -> Struct
// Map -> Struct
func Decode(input any, output any) error {
	return mapstructure.Decode(input, output)
}

func GetKeys(m map[string]any) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}

// Merge support multi Map, priority from low to high
func Merge(mapData ...map[string]any) (mapMerge map[string]any) {
	mapMerge = make(map[string]any)
	for _, mapItem := range mapData {
		for k, v := range mapItem {
			mapMerge[k] = v
		}
	}

	return
}

// New creates and returns an empty hash map.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
func New(safe ...bool) *gmap.Map {
	return gmap.New(safe...)
}

// SortKey sort by key
func SortKey(m map[string]any) (keys []string) {
	keys = GetKeys(m)
	sort.Strings(keys)

	return
}
