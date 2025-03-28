package vcmd_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestPortInUse(t *testing.T) {
	used, err := vcmd.PortInUse(90000)

	vtest.NotNil(t, err)
	vtest.Equal(t, false, used)
}
