package vmysql

import (
	"net/url"

	"gorm.io/driver/mysql"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vsafe"
	"github.com/sunshinexcode/gotk/vstr"
)

type (
	Mysql struct {
		Client  *DB
		Options *Options
	}

	Options struct {
		Charset  string `mapstructure:",omitempty"`
		Db       string `mapstructure:",omitempty"`
		Host     string `mapstructure:",omitempty"`
		Loc      string `mapstructure:",omitempty"`
		Password string `mapstructure:",omitempty"`
		Port     int    `mapstructure:",omitempty"`
		Timeout  string `mapstructure:",omitempty"`
		UserName string `mapstructure:",omitempty"`
	}
)

var (
	m *Mysql

	defaultOptions = map[string]any{
		"Charset":  "utf8",
		"Db":       "test",
		"Host":     "localhost",
		"Loc":      "Local",
		"Password": "",
		"Port":     3306,
		"Timeout":  "10s",
		"UserName": "root",
	}
)

// New create new mysql
func New(options map[string]any) (mysqlS *Mysql, err error) {
	mysqlS = &Mysql{Options: &Options{}}
	err = mysqlS.SetConfig(options)

	return
}

func Get() *Mysql {
	return m
}

func Set(mysql *Mysql) {
	m = mysql
}

// GetDsn get dsn
func (mysqlS *Mysql) GetDsn() string {
	return vstr.S("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%s", mysqlS.Options.UserName, mysqlS.Options.Password, mysqlS.Options.Host, mysqlS.Options.Port, mysqlS.Options.Db, mysqlS.Options.Charset, url.QueryEscape(mysqlS.Options.Loc), mysqlS.Options.Timeout)
}

// SetConfig set config
func (mysqlS *Mysql) SetConfig(options map[string]any) (err error) {
	if err = vreflect.SetAttrs(mysqlS.Options, vmap.Merge(defaultOptions, options)); err != nil {
		return
	}

	mysqlS.Client, err = Open(mysql.Open(mysqlS.GetDsn()), &Config{})
	return
}

func (o *Options) String() (data string) {
	data, _ = vjson.Encode(o)

	return vsafe.MaskPassword(data)
}
