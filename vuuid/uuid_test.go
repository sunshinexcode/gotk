package vuuid_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vuuid"
)

func TestGet(t *testing.T) {
	vtest.Equal(t, 36, len(vuuid.Get()))
}
