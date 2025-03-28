package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vmysql"

	"app/configs"
)

// NewMysql new mysql
func NewMysql(config *configs.Config) (mysql *vmysql.Mysql, err error) {
	return vbootstrap.NewMysql(config)
}
