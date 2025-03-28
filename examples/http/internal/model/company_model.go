package model

import (
	"context"
	"errors"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmodel"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtime"

	"app/internal/entity"
)

var _ ICompanyModel = (*CompanyModel)(nil)

type ICompanyModel interface {
	Create(ctx context.Context, CompanyReq *entity.Company) error
	GetIdsByPage(ctx context.Context, pageSize int, companyEntity *entity.Company) ([]*entity.Company, error)
	QueryById(ctx context.Context, id int64) (data *entity.Company, err error)
	UpdateById(ctx context.Context, id int64, companyReq *entity.Company) error
	UpdateByIdOfMap(ctx context.Context, id int64, data map[string]any) error
}

type CompanyModel struct {
	BaseModel *vmodel.BaseModel
}

func NewCompanyModel(baseModel *vmodel.BaseModel) *CompanyModel {
	return &CompanyModel{BaseModel: baseModel}
}

func (model *CompanyModel) create(ctx context.Context, companyEntity *entity.Company) error {
	companyEntity.CreateTime = vtime.GetNowUtc()
	companyEntity.UpdateTime = vtime.GetNowUtc()

	result := model.BaseModel.GetDb(ctx).Create(companyEntity)

	vlog.Infoc(ctx, "create", "err", result.Error, "companyEntity", companyEntity)
	return verror.Wrap(result.Error, vcode.CodeErrDbOperation, companyEntity)
}

func (model *CompanyModel) Create(ctx context.Context, companyReq *entity.Company) error {
	companyEntity := &entity.Company{
		CompanyName: companyReq.CompanyName,
	}

	vlog.Infoc(ctx, "Create", "companyEntity", companyEntity)
	return model.create(ctx, companyEntity)
}

func (model *CompanyModel) GetIdsByPage(ctx context.Context, pageSize int, companyEntity *entity.Company) ([]*entity.Company, error) {
	var list []*entity.Company

	result := model.BaseModel.Mysql.Client.Select("id, company_name").Where("id > ?", companyEntity.Id).Order("id ASC").Limit(pageSize).Find(&list)

	return list, verror.Wrap(result.Error, vcode.CodeErrDbOperation, result, list)
}

func (model *CompanyModel) QueryById(ctx context.Context, id int64) (data *entity.Company, err error) {
	data = &entity.Company{}

	result := model.BaseModel.Mysql.Client.Where("id = ?", id).First(data)

	if errors.Is(result.Error, vmysql.ErrRecordNotFound) {
		return data, verror.ErrDbRecordNotFound
	}

	return data, verror.Wrap(result.Error, vcode.CodeErrDbOperation, result, id, data)
}

func (model *CompanyModel) UpdateById(ctx context.Context, id int64, companyReq *entity.Company) error {
	data := map[string]any{
		"company_name": companyReq.CompanyName,
	}

	vlog.Infoc(ctx, "UpdateById", "id", id, "data", data)
	return model.UpdateByIdOfMap(ctx, id, data)
}

func (model *CompanyModel) UpdateByIdOfMap(ctx context.Context, id int64, data map[string]any) error {
	data["update_time"] = vtime.GetNowUtc()

	result := model.BaseModel.GetDb(ctx).Model(&entity.Company{}).Where("id = ?", id).Updates(data)

	vlog.Infoc(ctx, "UpdateByIdOfMap", "err", result.Error, "id", id, "data", data)
	return verror.Wrap(result.Error, vcode.CodeErrDbOperation, result, data)
}
