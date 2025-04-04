package vaes_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vaes"
	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vmd5"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestEncryptCbc(t *testing.T) {
	data, err := vaes.EncryptCbc([]byte("test"), vmd5.GetByte("test"), vmd5.GetByte("test"), vaes.Pkcs5Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "lGJ+/zteke3vdHNdAuVGMw==", vbase64.EncodeToStr(data))

	dataDecrypt, err := vaes.DecryptCbc(data, vmd5.GetByte("test"), vmd5.GetByte("test"), vaes.Pkcs5Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(dataDecrypt))

	data, err = vaes.EncryptCbc([]byte("test"), vmd5.GetByte("test"), vmd5.GetByte("test"), vaes.Pkcs7Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "lGJ+/zteke3vdHNdAuVGMw==", vbase64.EncodeToStr(data))

	dataDecrypt, err = vaes.DecryptCbc(data, vmd5.GetByte("test"), vmd5.GetByte("test"), vaes.Pkcs7Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(dataDecrypt))
}
