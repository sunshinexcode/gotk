package vrand_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vrand"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestGetNum(t *testing.T) {
	vtest.Equal(t, 1, vrand.GetNum(1, 1))
}

func TestGetStr(t *testing.T) {
	str, err := vrand.GetStr()

	vtest.Nil(t, err)
	vtest.Equal(t, 21, len(str))

	str, err = vrand.GetStr(5)

	vtest.Nil(t, err)
	vtest.Equal(t, 5, len(str))

	str, err = vrand.GetStr(10)

	vtest.Nil(t, err)
	vtest.Equal(t, 10, len(str))
}
