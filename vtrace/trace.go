package vtrace

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vuuid"
)

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	TraceIdKey = "traceId"
)

// GenerateTraceId return traceId
// creates and returns a global unique string in 32 bytes that meets most common
// usages without strict UUID algorithm. It returns a unique string using default
// unique algorithm if no `data` is given.
func GenerateTraceId() string {
	return guid.S()
}

// GetGinContextTraceId get trace id from context
func GetGinContextTraceId(c *gin.Context) string {
	if traceId, existed := c.Get(TraceIdKey); existed {
		return traceId.(string)
	}

	return ""
}

func GetTraceId(ctx context.Context) string {
	if _, ok := ctx.(*gin.Context); ok {
		return GetGinContextTraceId(ctx.(*gin.Context))
	}

	if val, ok := ctx.Value(ContextKey(TraceIdKey)).(string); ok {
		return val
	}

	return ""
}

func GetTraceIdKey() string {
	return TraceIdKey
}

func SetTraceId(ctx context.Context, prefix string, traceId ...string) context.Context {
	tid := vuuid.Get()

	if len(traceId) > 0 {
		tid = traceId[0]
	}

	return context.WithValue(ctx, ContextKey(TraceIdKey), vstr.S("%s%s", prefix, tid))
}

func SetTraceIdKey(traceIdKey string) {
	TraceIdKey = traceIdKey
}
