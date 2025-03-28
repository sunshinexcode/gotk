package vtrace_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtrace"
)

func TestContextKey(t *testing.T) {
	vtest.Equal(t, "test", vtrace.ContextKey("test").String())
}

func TestGenerateTraceId(t *testing.T) {
	vtest.Equal(t, 32, len(vtrace.GenerateTraceId()))
}

func TestGetGinContextTraceId(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, c.Err())
	vtest.Equal(t, "", vtrace.GetGinContextTraceId(c))

	c.Set(vtrace.TraceIdKey, "test")

	vtest.Equal(t, "test", vtrace.GetGinContextTraceId(c))
}

func TestGetTraceId(t *testing.T) {
	ctx := context.TODO()

	vtest.Equal(t, "", vtrace.GetTraceId(ctx))

	ctx = context.WithValue(ctx, vtrace.ContextKey(vtrace.TraceIdKey), "test")

	vtest.Equal(t, "test", vtrace.GetTraceId(ctx))
}

func TestGetTraceIdGinContext(t *testing.T) {
	ctx, _ := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, ctx.Err())
	vtest.Equal(t, "", vtrace.GetTraceId(ctx))

	ctx.Set(vtrace.TraceIdKey, "test")

	vtest.Equal(t, "test", vtrace.GetTraceId(ctx))
}

func TestGetTraceIdKey(t *testing.T) {
	vtest.Equal(t, "traceId", vtrace.GetTraceIdKey())
}

func TestSetTraceId(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "test-")

	vlog.Debug("TestSetTraceId", "vtrace.GetTraceId(ctx)", vtrace.GetTraceId(ctx))

	vtest.Equal(t, true, vtrace.GetTraceId(ctx) != "")

	ctx = vtrace.SetTraceId(context.TODO(), "test-", "test")

	vtest.Equal(t, "test-test", vtrace.GetTraceId(ctx))
}

func TestSetTraceIdKey(t *testing.T) {
	vtest.Equal(t, "traceId", vtrace.GetTraceIdKey())

	vtrace.SetTraceIdKey("requestId")

	vtest.Equal(t, "requestId", vtrace.GetTraceIdKey())
}
