package vbase64_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestDecode(t *testing.T) {
	data, err := vbase64.Decode([]byte("SGVsbG8gV29ybGQ="))

	vtest.Nil(t, err)
	vtest.Equal(t, "Hello World", string(data))
}

func TestDecodeStr(t *testing.T) {
	data, err := vbase64.DecodeStr("SGVsbG8gV29ybGQ=")

	vtest.Nil(t, err)
	vtest.Equal(t, "Hello World", string(data))
}

func TestDecodeToStr(t *testing.T) {
	data, err := vbase64.DecodeToStr("SGVsbG8gV29ybGQ=")

	vtest.Nil(t, err)
	vtest.Equal(t, "Hello World", data)
}

func TestEncode(t *testing.T) {
	vtest.Equal(t, "SGVsbG8gV29ybGQ=", string(vbase64.Encode([]byte("Hello World"))))
	vtest.Equal(t, "", string(vbase64.Encode([]byte(""))))
}

func TestEncodeStr(t *testing.T) {
	vtest.Equal(t, "SGVsbG8gV29ybGQ=", vbase64.EncodeStr("Hello World"))
	vtest.Equal(t, "", vbase64.EncodeStr(""))
}

func TestEncodeToStr(t *testing.T) {
	vtest.Equal(t, "SGVsbG8gV29ybGQ=", vbase64.EncodeToStr([]byte("Hello World")))
	vtest.Equal(t, "", vbase64.EncodeToStr([]byte("")))
}
