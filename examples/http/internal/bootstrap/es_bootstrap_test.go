package bootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/ves"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/bootstrap"
)

func TestNewEsMock(t *testing.T) {
	vconfig.SetConfigPath("../../configs")

	config, err := configs.New()

	vtest.Nil(t, err)

	patch := vmock.ApplyFuncReturn(vbootstrap.NewEs, &ves.Es{}, nil)
	defer vmock.Reset(patch)

	_, err = bootstrap.NewEs(config)

	vtest.Nil(t, err)
}
