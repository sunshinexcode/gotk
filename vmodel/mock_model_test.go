package vmodel_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmodel"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMock(t *testing.T) {
	vtest.Equal(t, true, len(vmodel.Mock(vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}}))) > 0)
}
