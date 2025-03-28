package venv_test

import (
	"os"
	"testing"

	"github.com/sunshinexcode/gotk/venv"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, "APP_ENV", venv.AppEnvKey)
	vtest.Equal(t, "dev", venv.AppEnvDev)
	vtest.Equal(t, "test", venv.AppEnvTest)
	vtest.Equal(t, "pre", venv.AppEnvPre)
	vtest.Equal(t, "prod", venv.AppEnvProd)
}

func TestGetEnv(t *testing.T) {
	vtest.Equal(t, "", venv.GetEnv("test", ""))
	vtest.Equal(t, "test", venv.GetEnv("test", "test"))

	os.Setenv("test", "val")

	vtest.Equal(t, "val", venv.GetEnv("test", ""))
}
