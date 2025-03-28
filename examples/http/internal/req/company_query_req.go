package req

import (
	"github.com/sunshinexcode/gotk/vreq"
)

type CompanyQueryReq struct {
	vreq.TraceIdReq

	Id int64 `form:"id,omitempty" json:"id,omitempty" v:"required|integer|min:1"`
}
