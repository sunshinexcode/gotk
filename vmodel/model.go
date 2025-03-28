package vmodel

import (
	"context"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vmysql"
	"github.com/sunshinexcode/gotk/vtrace"
)

type BaseModel struct {
	Mysql *vmysql.Mysql
}

const (
	TxContextKey = "dbTx"
)

func NewBaseModel(mysql *vmysql.Mysql) *BaseModel {
	return &BaseModel{Mysql: mysql}
}

func (model *BaseModel) Begin(ctx context.Context) (context.Context, *vmysql.DB) {
	tx := model.Mysql.Client.Begin()
	ctx = SetTx(ctx, tx)

	return ctx, tx
}

func (model *BaseModel) Commit() *vmysql.DB {
	return model.Mysql.Client.Commit()
}

func (model *BaseModel) GetDb(ctx context.Context) (db *vmysql.DB) {
	db = model.Mysql.Client

	if tx := GetTx(ctx); tx != nil {
		db = tx
	}

	return
}

func (model *BaseModel) Rollback() *vmysql.DB {
	return model.Mysql.Client.Rollback()
}

func GetTx(ctx context.Context) *vmysql.DB {
	if _, ok := ctx.(*vapi.Context); ok {
		if tx, existed := ctx.(*vapi.Context).Get(TxContextKey); existed {
			return tx.(*vmysql.DB)
		}
	}

	if tx, ok := ctx.Value(vtrace.ContextKey(TxContextKey)).(*vmysql.DB); ok {
		return tx
	}

	return nil
}

func SetTx(ctx context.Context, tx *vmysql.DB) context.Context {
	if _, ok := ctx.(*vapi.Context); ok {
		ctx.(*vapi.Context).Set(TxContextKey, tx)
		return ctx
	}

	return context.WithValue(ctx, vtrace.ContextKey(TxContextKey), tx)
}
