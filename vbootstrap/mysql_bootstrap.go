package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmysql"
)

// NewMysql new mysql
func NewMysql(config vconfig.IConfig) (mysql *vmysql.Mysql, err error) {
	var options vmap.M
	mysql = &vmysql.Mysql{}

	defer func() {
		if err != nil {
			vlog.Error("NewMysql", "err", err, "options", mysql.Options)
		}
	}()

	if err = vmap.Decode(config.GetMysql(), &options); err != nil {
		return
	}
	if mysql, err = vmysql.New(options); err != nil {
		return
	}

	vlog.Infof("NewMysql, options:%+v", mysql.Options)
	return
}
