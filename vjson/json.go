package vjson

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// Decode String => Object
func Decode(data string) (v map[string]any, err error) {
	err = json.Unmarshal([]byte(data), &v)
	return
}

// Encode Object => String
func Encode(v any) (data string, err error) {
	dataByte, err := json.Marshal(v)
	data = string(dataByte)

	return
}

// Parse parses the json and returns a result.
func Parse(json string) gjson.Result {
	return gjson.Parse(json)
}
