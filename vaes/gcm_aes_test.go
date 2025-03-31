package vaes_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vaes"
	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vmd5"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestDecryptGcm(t *testing.T) {
	_, err := vaes.DecryptGcm([]byte("test"), vmd5.GetByte("test"), vmd5.GetByte("test")[:12])

	vtest.NotNil(t, err)
	vtest.Equal(t, "cipher: message authentication failed", err.Error())

	_, err = vaes.DecryptGcm([]byte("test"), vmd5.GetByte("test")[:1], vmd5.GetByte("test")[:12])

	vtest.NotNil(t, err)
	vtest.Equal(t, "crypto/aes: invalid key size 1", err.Error())
}

func TestEncryptGcm(t *testing.T) {
	dst, err := vaes.EncryptGcm([]byte("test"), vmd5.GetByte("test"), vmd5.GetByte("test")[:12])

	vtest.Nil(t, err)
	vtest.Equal(t, "IBXU2jtsmYZBNoxnjINuTHreR0k=", vbase64.EncodeToStr(dst))

	src, err := vaes.DecryptGcm(dst, vmd5.GetByte("test"), vmd5.GetByte("test")[:12])

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(src))

	_, err = vaes.EncryptGcm([]byte("test"), vmd5.GetByte("test")[:1], vmd5.GetByte("test")[:12])

	vtest.NotNil(t, err)
	vtest.Equal(t, "crypto/aes: invalid key size 1", err.Error())
}
