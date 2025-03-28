package vmiddleware

import (
	"github.com/gin-gonic/gin/binding"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vtrace"
)

type (
	TraceId struct {
		GetTraceIdFunc    GetTraceIdFunc
		TraceIdContextKey string
		TraceIdKey        string
	}

	GetTraceIdFunc func(c *vapi.Context, traceId *TraceId, traceIdReq string) string
)

var (
	DefaultTraceId    = NewTraceId()
	TraceIdContextKey = vtrace.GetTraceIdKey()
)

// NewTraceId create traceId object
func NewTraceId() (t *TraceId) {
	t = &TraceId{
		GetTraceIdFunc:    GetTraceIdVal,
		TraceIdContextKey: TraceIdContextKey,
		TraceIdKey:        TraceIdContextKey,
	}
	return
}

// SetGetTraceIdFunc set GetTraceIdFunc
func (t *TraceId) SetGetTraceIdFunc(getTraceIdFunc GetTraceIdFunc) *TraceId {
	t.GetTraceIdFunc = getTraceIdFunc
	return t
}

// SetTraceIdKey set traceIdKey
func (t *TraceId) SetTraceIdKey(traceIdKey string) *TraceId {
	t.TraceIdKey = traceIdKey
	return t
}

// GetTraceIdVal get trace id
func GetTraceIdVal(c *vapi.Context, t *TraceId, traceIdReq string) (traceId string) {
	var existed bool

	if traceId, existed = c.GetQuery(t.TraceIdKey); existed {
		return
	}
	if traceId, existed = c.GetPostForm(t.TraceIdKey); existed {
		return
	}

	return traceIdReq
}

// GetTraceIdValByHeader get trace id by header
func GetTraceIdValByHeader(c *vapi.Context, t *TraceId, _ string) string {
	return c.GetHeader(t.TraceIdKey)
}

// TraceIdMiddleware trace id
func TraceIdMiddleware[T any](t *TraceId) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		var r T
		var traceIdReq string

		// Get traceId
		if c.GetHeader("Content-Type") == vapi.MimeJson {
			_ = c.ShouldBindBodyWith(&r, binding.JSON)
		} else {
			_ = c.Bind(&r)
		}

		m := vconv.Map(r)
		if c.GetHeader("Content-Type") == vapi.MimeJson {
			if traceIdM, ok := m[t.TraceIdKey].(string); ok {
				traceIdReq = traceIdM
			}
		}

		c.Set(t.TraceIdContextKey, t.GetTraceIdFunc(c, t, traceIdReq))
		c.Next()
	}
}
