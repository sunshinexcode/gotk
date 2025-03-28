package vmodel

import (
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmysql"
)

func Mock(m *BaseModel) (patches []*vmock.Patches) {
	patchModel := vmock.ApplyMethodReturn(m.Mysql.Client, "Model", &vmysql.DB{})
	patches = append(patches, patchModel)

	patchRaw := vmock.ApplyMethodReturn(m.Mysql.Client, "Raw", &vmysql.DB{})
	patches = append(patches, patchRaw)

	patchSelect := vmock.ApplyMethodReturn(m.Mysql.Client, "Select", &vmysql.DB{})
	patches = append(patches, patchSelect)

	patchWhere := vmock.ApplyMethodReturn(m.Mysql.Client, "Where", &vmysql.DB{})
	patches = append(patches, patchWhere)

	patchOrder := vmock.ApplyMethodReturn(m.Mysql.Client, "Order", &vmysql.DB{})
	patches = append(patches, patchOrder)

	patchLimit := vmock.ApplyMethodReturn(m.Mysql.Client, "Limit", &vmysql.DB{})
	patches = append(patches, patchLimit)

	return
}
