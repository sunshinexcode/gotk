package vjson

import (
	"encoding/json"
	"os"
)

func ConvertFileToStruct[T any](jsonFile string, obj *T) (err error) {
	jsonByte, err := os.ReadFile(jsonFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, &obj)
	return
}
