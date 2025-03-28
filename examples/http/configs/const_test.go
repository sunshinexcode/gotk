package configs_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, false, configs.TestOpen)
}
