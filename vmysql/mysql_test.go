package vmysql_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewMock(t *testing.T) {
	var err error

	mysql := &vmysql.Mysql{Options: &vmysql.Options{}}
	patch := vmock.ApplyMethod(reflect.TypeOf(mysql), "SetConfig", func(mysql *vmysql.Mysql, options map[string]any) error {
		mysql.Options = &vmysql.Options{Host: "localhost", Password: "123"}
		return nil
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, "", mysql.Options.Host)

	mysql, err = vmysql.New(nil)
	vmysql.Set(mysql)

	vtest.Nil(t, err)
	vtest.Equal(t, "localhost", mysql.Options.Host)

	mysql2 := vmysql.Get()

	vtest.Equal(t, "localhost", mysql2.Options.Host)
}

func TestGetDsn(t *testing.T) {
	mysql := &vmysql.Mysql{Options: &vmysql.Options{}}

	vtest.Equal(t, "", mysql.Options.Host)
	vtest.Equal(t, ":@tcp(:0)/?charset=&parseTime=True&loc=&timeout=", mysql.GetDsn())

	mysql = &vmysql.Mysql{Options: &vmysql.Options{UserName: "test"}}

	vtest.Equal(t, "test:@tcp(:0)/?charset=&parseTime=True&loc=&timeout=", mysql.GetDsn())
}

func TestSetConfig(t *testing.T) {
	mysql := &vmysql.Mysql{Options: &vmysql.Options{}}
	err := mysql.SetConfig(vmap.M{"Port": 1})

	vtest.NotNil(t, err)
	vtest.Equal(t, "localhost", mysql.Options.Host)
	vtest.Equal(t, true, strings.Contains(err.Error(), "connect: connection refused"))

	err = mysql.SetConfig(vmap.M{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
}

func TestString(t *testing.T) {
	vtest.Equal(t, `{"Charset":"","Db":"","Host":"","Loc":"","Password":"***"}`, vstr.S("%+v", &vmysql.Options{Password: "123"}))
}
