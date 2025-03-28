package model_test

import (
	"context"
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmodel"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtest"

	"app/internal/entity"
	"app/internal/model"
)

func initCompanyModel() (m *model.CompanyModel, patches []*vmock.Patches) {
	m = model.NewCompanyModel(vmodel.NewBaseModel(&vmysql.Mysql{Client: &vmysql.DB{}}))
	patches = vmodel.Mock(m.BaseModel)

	return
}

func TestCompanyModelCreateMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Create", &vmysql.DB{})
	defer vmock.Reset(patch)

	vtest.Nil(t, m.Create(context.TODO(), &entity.Company{}))
}

func TestCompanyModelCreateErrorMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Create", &vmysql.DB{Error: errors.New("create error")})
	defer vmock.Reset(patch)

	err := m.Create(context.TODO(), &entity.Company{})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, errors.Is(err, verror.ErrDbOperation))
}

func TestCompanyModelGetIdsByPageMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Find", &vmysql.DB{})
	defer vmock.Reset(patch)

	list, err := m.GetIdsByPage(context.TODO(), 10, &entity.Company{})

	vtest.Nil(t, err)
	vtest.Equal(t, 0, len(list))
}

func TestCompanyModelGetIdsByPageErrorMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Find", &vmysql.DB{Error: errors.New("find error")})
	defer vmock.Reset(patch)

	list, err := m.GetIdsByPage(context.TODO(), 10, &entity.Company{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10120|database operation error|&{Config:<nil> Error:find error RowsAffected:0 Statement:<nil> clone:0} []  \n-> find error", err.Error())
	vtest.Equal(t, 0, len(list))
}

func TestCompanyModelQueryByIdMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodFunc(m.BaseModel.Mysql.Client, "First", func(dest interface{}, conds ...interface{}) (tx *vmysql.DB) {
		dest.(*entity.Company).Id = 1
		return &vmysql.DB{}
	})
	defer vmock.Reset(patch)

	data, err := m.QueryById(context.TODO(), 1)

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), data.Id)
}

func TestCompanyModelQueryByIdErrorMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "First", &vmysql.DB{Error: vmysql.ErrRecordNotFound})
	defer vmock.Reset(patch)

	data, err := m.QueryById(context.TODO(), 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "10121|database record not found|<nil>", err.Error())
	vtest.Equal(t, int64(0), data.Id)
}

func TestCompanyModelUpdateByIdMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Updates", &vmysql.DB{})
	defer vmock.Reset(patch)

	vtest.Nil(t, m.UpdateById(context.TODO(), 1, &entity.Company{}))
}

func TestCompanyModelUpdateByIdErrorMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Updates", &vmysql.DB{Error: errors.New("update error")})
	defer vmock.Reset(patch)

	err := m.UpdateById(context.TODO(), 1, &entity.Company{})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, errors.Is(err, verror.ErrDbOperation))
}

func TestCompanyModelUpdateByIdOfMapMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Updates", &vmysql.DB{})
	defer vmock.Reset(patch)

	vtest.Nil(t, m.UpdateByIdOfMap(context.TODO(), 1, vmap.M{}))
}

func TestCompanyModelUpdateByIdOfMapErrorMock(t *testing.T) {
	m, patches := initCompanyModel()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(m.BaseModel.Mysql.Client, "Updates", &vmysql.DB{Error: errors.New("update error")})
	defer vmock.Reset(patch)

	err := m.UpdateByIdOfMap(context.TODO(), 1, vmap.M{})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, errors.Is(err, verror.ErrDbOperation))
}
