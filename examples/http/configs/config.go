package configs

import (
	"embed"

	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
)

type (
	Config struct {
		vconfig.Config `mapstructure:",squash"`

		AppCustom *AppCustom
	}

	AppCustom struct {
		ConsoleThirdPartyHost          string
		ConsoleThirdPartyAuthorization string

		CronPatternCreateTaskForVerify  string
		CronPatternProcessTaskForVerify string
	}
)

//go:embed config-default.toml config-dev.toml config-pre.toml config-prod.toml config-test.toml
var fs embed.FS

func New() (*Config, error) {
	InitConfig()
	return NewConfig()
}

func NewConfig() (*Config, error) {
	config, err := vconfig.NewConfig(&Config{})
	if err != nil {
		vlog.Fatalf("%s", err)
	}

	vlog.Infof("NewConfig, config:%+v", config)
	return config, err
}

func InitConfig() {
	if err := vconfig.InitConfigEmbed(fs); err != nil {
		vlog.Fatalf("%s", err)
	}
}

func GetConfig() (config *Config) {
	InitConfig()
	config, _ = NewConfig()
	return
}
