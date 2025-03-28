package vmd5_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vmd5"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestGet(t *testing.T) {
	md5, err := vmd5.Get("test")

	vtest.Nil(t, err)
	vtest.Equal(t, "098f6bcd4621d373cade4e832627b4f6", md5)
	vtest.Equal(t, 32, len(md5))
}

func TestGetByte(t *testing.T) {
	md5 := vmd5.GetByte("test")

	vtest.Equal(t, "CY9rzUYh03PK3k6DJie09g==", vbase64.EncodeToStr(vmd5.GetByte("test")))
	vtest.Equal(t, 16, len(md5))
}

func TestGetStr(t *testing.T) {
	md5 := vmd5.GetStr("test")

	vtest.Equal(t, "098f6bcd4621d373cade4e832627b4f6", md5)
	vtest.Equal(t, 32, len(md5))
}
