package vvar_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vvar"
)

func TestNew(t *testing.T) {
	vtest.Equal(t, "test", vvar.New("test").String())
}
