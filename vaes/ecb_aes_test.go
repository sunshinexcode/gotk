package vaes_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vaes"
	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vmd5"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestEncryptECB(t *testing.T) {
	data, err := vaes.EncryptECB([]byte("test"), vmd5.GetByte("test"), vaes.Pkcs5Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "Neenfr7yTXFYG8T2p/t38A==", vbase64.EncodeToStr(data))

	dataDecrypt, err := vaes.DecryptECB(data, vmd5.GetByte("test"), vaes.Pkcs5Padding)

	vtest.Nil(t, err)
	vtest.Equal(t, "test", string(dataDecrypt))
}
