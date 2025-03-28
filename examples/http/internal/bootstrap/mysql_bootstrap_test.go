package bootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/bootstrap"
)

func TestNewMysqlMock(t *testing.T) {
	vconfig.SetConfigPath("../../configs")

	config, err := configs.New()

	vtest.Nil(t, err)

	patch := vmock.ApplyFuncReturn(vbootstrap.NewMysql, &vmysql.Mysql{}, nil)
	defer vmock.Reset(patch)

	_, err = bootstrap.NewMysql(config)

	vtest.Nil(t, err)
}
