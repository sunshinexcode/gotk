package vconfig_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestUnmarshal(t *testing.T) {
	vtest.Nil(t, vconfig.Unmarshal(&vmap.M{}))
}
