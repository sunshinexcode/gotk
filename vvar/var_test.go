package vvar_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vvar"
)

func TestIsNil(t *testing.T) {
	vtest.Equal(t, true, vvar.IsNil(nil))

	intValue := 1

	vtest.Equal(t, false, vvar.IsNil(&intValue))
	vtest.Equal(t, false, vvar.IsNil(intValue))
	vtest.Equal(t, false, vvar.IsNil(2))
	vtest.Equal(t, false, vvar.IsNil(map[string]string{}))
	vtest.Equal(t, false, vvar.IsNil([]string{}))
	vtest.Equal(t, true, vvar.IsNil(interface{}(nil)))
	vtest.Equal(t, false, vvar.IsNil(interface{}("a")))
}
