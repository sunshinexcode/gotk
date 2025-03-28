package vjson_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vtest"
)

type CompanyQueryResp struct {
	Name string `json:"name,omitempty"`
}

func TestConvertJsonFileToStruct(t *testing.T) {
	data := &CompanyQueryResp{}
	err := vjson.ConvertFileToStruct("../examples/http/data/testing/company_query.json", data)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", data.Name)
}

func TestConvertJsonFileToStructError(t *testing.T) {
	data := &CompanyQueryResp{}
	err := vjson.ConvertFileToStruct("../examples/http/data/testing/file_not_existed.json", data)

	vtest.NotNil(t, err)
	vtest.Equal(t, "open ../examples/http/data/testing/file_not_existed.json: no such file or directory", err.Error())
	vtest.Equal(t, "", data.Name)
}
