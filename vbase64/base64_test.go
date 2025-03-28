package vbase64_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestDecode(t *testing.T) {
	data, err := vbase64.Decode([]byte("dGVzdA=="))

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(data))
}

func TestDecodeStr(t *testing.T) {
	data, err := vbase64.DecodeStr("dGVzdA==")

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(data))
}

func TestDecodeToStr(t *testing.T) {
	data, err := vbase64.DecodeToStr("dGVzdA==")

	vtest.Nil(t, err)
	vtest.Equal(t, "test", data)
}

func TestEncode(t *testing.T) {
	vtest.Equal(t, "dGVzdA==", string(vbase64.Encode([]byte("test"))))
	vtest.Equal(t, "", string(vbase64.Encode([]byte(""))))
}

func TestEncodeStr(t *testing.T) {
	vtest.Equal(t, "dGVzdA==", vbase64.EncodeStr("test"))
	vtest.Equal(t, "", vbase64.EncodeStr(""))
}

func TestEncodeToStr(t *testing.T) {
	vtest.Equal(t, "dGVzdA==", vbase64.EncodeToStr([]byte("test")))
	vtest.Equal(t, "", vbase64.EncodeToStr([]byte("")))
}
