package vmiddleware_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtrace"
)

func TestTraceIdMiddleware(t *testing.T) {
	traceId := vtrace.GenerateTraceId()
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[Request](vmiddleware.DefaultTraceId), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.DefaultTraceId.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)

		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test?name=test1&age=18&traceId=%s", traceId), nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestTraceIdMiddlewarePost(t *testing.T) {
	traceId := vtrace.GenerateTraceId()
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[Request](vmiddleware.DefaultTraceId), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.DefaultTraceId.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)

		voutput.O(c, nil, "test")
	})

	req := httptest.NewRequest(http.MethodPost, "/test?name=test2", bytes.NewBuffer([]byte(fmt.Sprintf("traceId=%s", traceId))))
	req.Header.Add("Content-Type", vapi.MimePostForm)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestTraceIdMiddlewarePostMimeJson(t *testing.T) {
	traceId := vtrace.GenerateTraceId()
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[RequestTraceId](vmiddleware.DefaultTraceId), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.DefaultTraceId.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)

		voutput.O(c, nil, "test")
	})

	m := vconv.Map(RequestTraceId{Name: "test", TraceId: traceId})
	dataStr, err := vjson.Encode(m)

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestTraceIdMiddlewarePostMimeJsonRequestNoTraceId(t *testing.T) {
	traceId := vtrace.GenerateTraceId()
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[RequestNoTraceId](vmiddleware.NewTraceId().SetTraceIdKey("X-TraceId").
		SetGetTraceIdFunc(vmiddleware.GetTraceIdValByHeader)), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.DefaultTraceId.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)

		voutput.O(c, nil, "test")
	})

	m := vconv.Map(RequestNoTraceId{Name: "test"})
	dataStr, err := vjson.Encode(m)

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	req.Header.Add("X-TraceId", traceId)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestTraceIdMiddlewarePostMimeJsonRequestRequestId(t *testing.T) {
	traceId := vtrace.GenerateTraceId()
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[RequestRequestId](vmiddleware.NewTraceId().SetTraceIdKey("requestId")), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.DefaultTraceId.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)
		vtest.Equal(t, traceId, vtrace.GetGinContextTraceId(c))

		voutput.O(c, nil, "test")
	})

	m := vconv.Map(RequestRequestId{Name: "test", RequestId: traceId})
	dataStr, err := vjson.Encode(m)

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestTraceIdMiddlewareSignMiddleware(t *testing.T) {
	traceId := "123456"
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[RequestTraceId](vmiddleware.NewTraceId()), vmiddleware.SignMiddleware[RequestTraceId](vmiddleware.NewSign().SetSecret(secret)), func(c *vapi.Context) {
		traceIdVal, _ := c.Get(vmiddleware.TraceIdContextKey)

		vtest.Equal(t, traceId, traceIdVal)

		voutput.O(c, nil, "test")
	})

	m := vmap.M{"name": "test", "age": 18, "traceId": traceId}
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "59150b25346ff77cdbc81402c5fe10b5", signVal)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test?name=test&age=18&traceId=%s&sign=%s", traceId, signVal), nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}
