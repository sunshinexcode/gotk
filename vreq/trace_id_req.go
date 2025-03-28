package vreq

type TraceIdReq struct {
	TraceId string `form:"traceId,omitempty" json:"traceId,omitempty" v:"required|length:1,100"`
}
