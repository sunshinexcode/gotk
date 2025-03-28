package vmodel_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmodel"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestBaseModelBegin(t *testing.T) {
	m := vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}})

	patch := vmock.ApplyMethodReturn(m.Mysql.Client, "Begin", &vmysql.DB{})
	defer vmock.Reset(patch)

	ctx, tx := m.Begin(context.TODO())

	vtest.Nil(t, tx.Error)
	vtest.Equal(t, true, vmodel.GetTx(ctx) != nil)
	vtest.Equal(t, true, vmodel.GetTx(context.TODO()) == nil)
}

func TestBaseModelBeginGinCtx(t *testing.T) {
	m := vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}})

	patch := vmock.ApplyMethodReturn(m.Mysql.Client, "Begin", &vmysql.DB{})
	defer vmock.Reset(patch)

	w := httptest.NewRecorder()
	c, _ := vapi.CreateTestContext(w)

	vtest.Equal(t, true, vmodel.GetTx(c) == (*vmysql.DB)(nil))
	vtest.Equal(t, true, vmodel.GetTx(c) == nil)

	c2, _ := vapi.CreateTestContext(w)
	ctx := vmodel.SetTx(c2, m.Mysql.Client.Begin())

	vtest.Equal(t, true, vmodel.GetTx(ctx) != (*vmysql.DB)(nil))
	vtest.Equal(t, true, vmodel.GetTx(ctx) != nil)

	c3, _ := vapi.CreateTestContext(w)
	ctx, db := m.Begin(c3)

	vtest.Nil(t, db.Error)
	vtest.Equal(t, true, vmodel.GetTx(ctx) != (*vmysql.DB)(nil))
	vtest.Equal(t, true, vmodel.GetTx(ctx) != nil)
}

func TestBaseModelCommit(t *testing.T) {
	m := vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}})

	patch := vmock.ApplyMethodReturn(m.Mysql.Client, "Commit", &vmysql.DB{})
	defer vmock.Reset(patch)

	vtest.Nil(t, m.Commit().Error)
}

func TestBaseModelGetDb(t *testing.T) {
	m := vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}})

	vtest.Nil(t, m.GetDb(context.TODO()).Error)
	vtest.Nil(t, m.GetDb(context.Background()).Error)

	patch := vmock.ApplyMethodReturn(m.Mysql.Client, "Begin", &vmysql.DB{})
	defer vmock.Reset(patch)

	ctx, _ := m.Begin(context.TODO())
	vtest.Nil(t, m.GetDb(ctx).Error)
}

func TestBaseModelRollback(t *testing.T) {
	m := vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}})

	patch := vmock.ApplyMethodReturn(m.Mysql.Client, "Rollback", &vmysql.DB{})
	defer vmock.Reset(patch)

	vtest.Nil(t, m.Rollback().Error)
}
