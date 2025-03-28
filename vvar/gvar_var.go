package vvar

import "github.com/gogf/gf/v2/container/gvar"

type (
	Var = gvar.Var
)

func New(value interface{}, safe ...bool) *Var {
	return gvar.New(value, safe...)
}
