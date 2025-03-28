package configs_test

import (
	"errors"
	"os"
	"testing"

	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/venv"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
)

func TestNew(t *testing.T) {
	vconfig.SetConfigPath("../configs")
	config, err := configs.New()

	vtest.Nil(t, err)
	vtest.Equal(t, "", os.Getenv(venv.AppEnvKey))
	vtest.Equal(t, "dev", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.DebugLevel, config.Log.Level)
	vtest.Equal(t, "debug", config.HttpServer.Model)
}

func TestNewConfig(t *testing.T) {
	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "dev", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.DebugLevel, config.Log.Level)
	vtest.Equal(t, "debug", config.HttpServer.Model)
}

func TestNewConfigTest(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvTest)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "test", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.InfoLevel, config.Log.Level)
	vtest.Equal(t, "test", config.HttpServer.Model)
}

func TestNewConfigPre(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvPre)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "pre", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.InfoLevel, config.Log.Level)
	vtest.Equal(t, "release", config.HttpServer.Model)
}

func TestNewConfigProd(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvProd)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "prod", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.InfoLevel, config.Log.Level)
	vtest.Equal(t, "release", config.HttpServer.Model)
}

func TestNewConfigErrorMock(t *testing.T) {
	patchUnmarshal := vmock.ApplyFuncReturn(vconfig.Unmarshal, errors.New("unmarshal error"))
	defer patchUnmarshal.Reset()

	patchFatalf := vmock.ApplyFuncReturn(vlog.Fatalf)
	defer patchFatalf.Reset()

	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.NotNil(t, err)
	vtest.Equal(t, "unmarshal error", err.Error())
	vtest.Equal(t, "", config.App.Env)
	vtest.Equal(t, "", config.HttpServer.Model)
}

func TestInitConfig(t *testing.T) {
	vtest.Nil(t, os.Setenv(venv.AppEnvKey, venv.AppEnvDev))

	vconfig.SetConfigPath("../configs")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "dev", config.App.Env)
	vtest.Equal(t, "/data/logs/app.log", config.Log.File)
	vtest.Equal(t, vlog.DebugLevel, config.Log.Level)
	vtest.Equal(t, "debug", config.HttpServer.Model)
}

func TestInitConfigErrorMock(t *testing.T) {
	patchFatalf := vmock.ApplyFuncReturn(vlog.Fatalf)
	defer patchFatalf.Reset()

	vtest.Nil(t, os.Setenv(venv.AppEnvKey, "not_existed"))

	vconfig.SetConfigPath("../configs-not-existed")
	configs.InitConfig()
	config, err := configs.NewConfig()

	vtest.Nil(t, err)
	vtest.Equal(t, "", config.App.Env)
}

func TestGetConfig(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvDev)

	vtest.Nil(t, err)

	config := configs.GetConfig()

	vtest.Equal(t, "dev", config.App.Env)
}

func TestSetConfigPath(t *testing.T) {
	vconfig.SetConfigPath("../configs")
}
