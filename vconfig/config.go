package vconfig

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"github.com/sunshinexcode/gotk/venv"
	"github.com/sunshinexcode/gotk/ves"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmongodb"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vsafe"
	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vwebsocket"
)

var (
	configPath = "configs"
	configType = "toml"
)

type (
	IConfig interface {
		GetEs() *ves.Options
		GetHttpServer() HttpServer
		GetLog() *vlog.Options
		GetMetric() *vmetric.Options
		GetMongodb() *vmongodb.Options
		GetMysql() *vmysql.Options
		GetRedis() *redis.UniversalOptions
		GetTcpServer() vtcp.Server
		GetWebsocketServer() vwebsocket.Server
	}

	Config struct {
		Api             Api
		App             App
		Es              *ves.Options
		HttpServer      HttpServer
		Log             *vlog.Options
		Metric          *vmetric.Options
		Mongodb         *vmongodb.Options
		Mysql           *vmysql.Options
		Redis           *redis.UniversalOptions
		TcpServer       vtcp.Server
		WebsocketServer vwebsocket.Server
	}

	Api struct {
		Secret string `mapstructure:",omitempty"`
	}

	App struct {
		Env  string `mapstructure:",omitempty"`
		Name string `mapstructure:",omitempty"`
	}

	HttpServer struct {
		Address string `mapstructure:",omitempty"`
		Model   string `mapstructure:",omitempty"`
	}
)

// NewConfig new config
func NewConfig[T any](conf *T) (config *T, err error) {
	if err = Unmarshal(conf); err != nil {
		vlog.Errorf("%s", err)
	}

	return conf, err
}

func (c *Config) GetEs() *ves.Options {
	return c.Es
}

func (c *Config) GetHttpServer() HttpServer {
	return c.HttpServer
}

func (c *Config) GetLog() *vlog.Options {
	return c.Log
}

func (c *Config) GetMetric() *vmetric.Options {
	return c.Metric
}

func (c *Config) GetMongodb() *vmongodb.Options {
	return c.Mongodb
}

func (c *Config) GetMysql() *vmysql.Options {
	return c.Mysql
}

func (c *Config) GetRedis() *redis.UniversalOptions {
	return c.Redis
}

func (c *Config) GetTcpServer() vtcp.Server {
	return c.TcpServer
}

func (c *Config) GetWebsocketServer() vwebsocket.Server {
	return c.WebsocketServer
}

// InitConfig init config
func InitConfig() (err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.SetConfigName("config-default")

	// Default
	if err = viper.ReadInConfig(); err != nil {
		vlog.Errorf("%s", err)
		return
	}

	configs := viper.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	// Env
	env := venv.GetEnv(venv.AppEnvKey, venv.AppEnvDev)
	viper.SetConfigName(fmt.Sprintf("config-%s", env))

	if err = viper.ReadInConfig(); err != nil {
		vlog.Errorf("%s", err)
		return
	}

	return
}

// InitConfigEmbed init config
func InitConfigEmbed(fs embed.FS) (err error) {
	viper.SetConfigType(configType)

	data, err := fs.ReadFile(fmt.Sprintf("config-default.%s", configType))
	if err != nil {
		vlog.Errorf("%s", err)
		return
	}

	// Default
	if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		vlog.Errorf("%s", err)
		return
	}

	configs := viper.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	// Env
	env := venv.GetEnv(venv.AppEnvKey, venv.AppEnvDev)
	data, err = fs.ReadFile(fmt.Sprintf("config-%s.%s", env, configType))
	if err != nil {
		vlog.Errorf("%s", err)
		return
	}

	if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		vlog.Errorf("%s", err)
		return
	}

	return
}

// SetConfigPath set config path
func SetConfigPath(confPath string) {
	configPath = confPath
}

// SetConfigType set config type
func SetConfigType(confType string) {
	configType = confType
}

func (a Api) String() (data string) {
	data, _ = vjson.Encode(a)
	return vsafe.MaskPassword(data)
}
