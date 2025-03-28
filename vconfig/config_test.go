package vconfig_test

import (
	"embed"
	"os"
	"testing"

	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/venv"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
)

//go:embed config-default.toml config-dev.toml
var fs embed.FS

func initConfig(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvDev)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.Nil(t, err)
}

func TestNewConfig(t *testing.T) {
	m := vmap.M{}
	_, err := vconfig.NewConfig(&m)

	vtest.Nil(t, err)

	var a = 100
	_, err = vconfig.NewConfig(&a)

	vtest.NotNil(t, err)
	vtest.Equal(t, `'' expected type 'int', got unconvertible type 'map[string]interface {}', value: 'map[]'`, err.Error())

	vconfig.SetConfigPath("./configs-not-existed")
	err = vconfig.InitConfig()

	vtest.NotNil(t, err)
	vtest.NotEqual(t, "", err.Error())

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "", config.App.Env)
	vtest.Nil(t, config.Log)

	err = os.Setenv(venv.AppEnvKey, "not_existed")

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.NotNil(t, err)

	config, err = vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.Log)
	vtest.Equal(t, "", config.App.Env)

	err = os.Setenv(venv.AppEnvKey, venv.AppEnvDev)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.Nil(t, err)

	config, err = vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.Log)
	vtest.Equal(t, "dev", config.App.Env)
}

func TestNewConfigTest(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvTest)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.Nil(t, err)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "test", config.App.Env)
}

func TestNewConfigPre(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvPre)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.Nil(t, err)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "pre", config.App.Env)
}

func TestGetEs(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetEs())
}

func TestGetHttpServer(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "", config.GetHttpServer().Address)
}

func TestGetLog(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetLog())
}

func TestGetMetric(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetMetric())
}

func TestGetMongodb(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetMongodb())
}

func TestGetMysql(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetMysql())
}

func TestGetRedis(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Nil(t, config.GetRedis())
}

func TestGetTcpServer(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "", config.GetTcpServer().Address)
}

func TestGetWebsocketServer(t *testing.T) {
	initConfig(t)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, int(0), config.GetWebsocketServer().Port)
}

func TestInitConfig(t *testing.T) {
	err := os.Setenv(venv.AppEnvKey, venv.AppEnvDev)

	vtest.Nil(t, err)

	vconfig.SetConfigPath("./")
	err = vconfig.InitConfig()

	vtest.Nil(t, err)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "dev", config.App.Env)
}

func TestInitConfigEmbed(t *testing.T) {
	err := vconfig.InitConfigEmbed(fs)

	vtest.Nil(t, err)

	config, err := vconfig.NewConfig(&vconfig.Config{})

	vtest.Nil(t, err)
	vtest.Equal(t, "dev", config.App.Env)

	err = os.Setenv(venv.AppEnvKey, venv.AppEnvTest)

	vtest.Nil(t, err)

	err = vconfig.InitConfigEmbed(fs)

	vtest.NotNil(t, err)
	vtest.Equal(t, "open config-test.toml: file does not exist", err.Error())

	err = os.Setenv(venv.AppEnvKey, venv.AppEnvDev)
	vconfig.SetConfigType("yaml")

	vtest.Nil(t, err)

	err = vconfig.InitConfigEmbed(fs)

	vtest.NotNil(t, err)
	vtest.Equal(t, "open config-default.yaml: file does not exist", err.Error())
}

func TestSetConfigPath(t *testing.T) {
	vconfig.SetConfigPath("./")
}

func TestString(t *testing.T) {
	vtest.Equal(t, `{"Secret":"***"}`, vstr.S("%s", vconfig.Api{Secret: "123"}))
	vtest.Equal(t, `{"Secret":"***"}`, vstr.S("%+v", vconfig.Api{Secret: "123"}))
}
