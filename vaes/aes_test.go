package vaes_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vaes"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, "PKCS5", vaes.Pkcs5Padding)
	vtest.Equal(t, "PKCS7", vaes.Pkcs7Padding)
	vtest.Equal(t, "ZEROS", vaes.ZerosPadding)
}
