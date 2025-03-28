package vjson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestDecode(t *testing.T) {
	data, err := vjson.Decode(`{"code": "0", "data": "test"}`)

	vtest.Equal(t, "0", data["code"])
	vtest.Nil(t, err)

	// error
	data, err = vjson.Decode(`{"code": "0", "data": 'test'}`)

	vtest.Nil(t, data)
	vtest.NotNil(t, err)
}

func TestEncode(t *testing.T) {
	data, err := vjson.Encode(map[string]any{"code": "0", "data": "test"})

	vtest.Equal(t, `{"code":"0","data":"test"}`, data)
	vtest.Nil(t, err)

	// error
	data, err = vjson.Encode(map[any]any{1: 0, "data": "test"})

	assert.Empty(t, data)
	vtest.NotNil(t, err)
}

func TestParse(t *testing.T) {
	vtest.Equal(t, "test", vjson.Parse(`{"name": "test"}`).Get("name").Str)
}
