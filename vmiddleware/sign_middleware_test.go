package vmiddleware_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmiddleware"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCalSign(t *testing.T) {
	sign, err := vmiddleware.CalSign(vmap.M{"name": "test1", "age": 18}, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "df2cd72723547bbc36478da624c67109", sign)

	sign, err = vmiddleware.CalSign(vmap.M{"name": "test1", "age": 0}, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "aeb574ad0302796736ffc74888546c97", sign)

	sign, err = vmiddleware.CalSign(vmap.M{"name": "test1", "age": 18}, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "1d5c81b59a44991620d34c26ff6b5063221ebcf891a9f474642b8f9ff71a693c", sign)

	sign, err = vmiddleware.CalSign(vmap.M{"name": "test1", "age": 0}, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "cc2cce2210555832b71fa02f82a319947fea6e135dc105b4d346ee0b03972206", sign)

	sign, err = vmiddleware.CalSign(vmap.M{"name": "test1"}, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "3c99360d6d78e4c104920e4e30319ad61990ef8d37ebe6e80a7fe4a59fe71d18", sign)

	m := vconv.Map(Request{Name: "test"})
	mStr, _ := vjson.Encode(m)

	vtest.Equal(t, `{"age":0,"name":"test","sign":""}`, mStr)

	sign, err = vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", sign)
}

func TestCalSignRequest(t *testing.T) {
	m := vconv.Map(Request{Name: "test"})
	mStr, _ := vjson.Encode(m)

	vtest.Equal(t, `{"age":0,"name":"test","sign":""}`, mStr)

	sign, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", sign)

	m = vconv.Map(Request{Name: "test", Age: 0})
	mStr, _ = vjson.Encode(m)

	vtest.Equal(t, `{"age":0,"name":"test","sign":""}`, mStr)

	sign, err = vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", sign)

	m = vconv.Map(Request{Name: "test", Age: 0, Sign: ""})
	mStr, _ = vjson.Encode(m)

	vtest.Equal(t, `{"age":0,"name":"test","sign":""}`, mStr)

	sign, err = vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", sign)
}

func TestCalSignRequestOmitempty(t *testing.T) {
	m := vconv.Map(RequestOmitempty{Name: "test"}, vconv.MapOption{OmitEmpty: true})
	mStr, _ := vjson.Encode(m)

	vtest.Equal(t, `{"name":"test"}`, mStr)

	sign, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "88cd2108b5347d973cf39cdf9053d7dd42704876d8c9a9bd8e2d168259d3ddf7", sign)

	m = vconv.Map(RequestOmitempty{Name: "test", Age: 0}, vconv.MapOption{OmitEmpty: true})
	mStr, _ = vjson.Encode(m)

	vtest.Equal(t, `{"name":"test"}`, mStr)

	sign, err = vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "88cd2108b5347d973cf39cdf9053d7dd42704876d8c9a9bd8e2d168259d3ddf7", sign)

	m = vconv.Map(RequestOmitempty{Name: "test", Age: 0, Sign: ""}, vconv.MapOption{OmitEmpty: true})
	mStr, _ = vjson.Encode(m)

	vtest.Equal(t, `{"name":"test"}`, mStr)

	sign, err = vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "88cd2108b5347d973cf39cdf9053d7dd42704876d8c9a9bd8e2d168259d3ddf7", sign)
}

func TestSignMiddleware(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.SignMiddleware[Request](vmiddleware.NewSign().SetSecret(secret)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	m := vmap.M{"name": "test", "age": 18}
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "fd5618174ab2a9328b66fd77ace3037f", signVal)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test?name=test&age=18&sign=%s", signVal), nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestSignMiddlewarePost(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.SignMiddleware[Request](vmiddleware.NewSign().SetSecret(secret)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	m := vmap.M{"name": "test", "age": 18}
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "fd5618174ab2a9328b66fd77ace3037f", signVal)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(fmt.Sprintf("name=test4&sign=%s", signVal))))
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusUnauthorized, w.Code)
	vtest.Equal(t, int64(vcode.CodeErrSignInvalid.Code()), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, vcode.CodeErrSignInvalid.Message(), vjson.Parse(w.Body.String()).Get("msg").Str)
}

func TestSignMiddlewarePostMimeJson(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.SignMiddleware[Request](vmiddleware.NewSign().SetSecret(secret).
		SetAlgorithm(vmiddleware.SignAlgorithmHmac).SetCalSignFunc(vmiddleware.CalSign)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	m := vconv.Map(Request{Name: "test"})
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", signVal)

	m["sign"] = signVal
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

func TestSignMiddlewarePostMimePostForm(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.SignMiddleware[Request](vmiddleware.NewSign().SetSecret(secret)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	m := vconv.Map(Request{Name: "test"})
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmMd5)

	vtest.Nil(t, err)
	vtest.Equal(t, "8d927f5673b53f63521ab1c5e3739eef", signVal)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(fmt.Sprintf("name=test&age=0&sign=%s", signVal))))
	req.Header.Add("Content-Type", vapi.MimePostForm)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestSignMiddlewareSetSignKey(t *testing.T) {
	g := vapi.Default()
	g.Any("/test", vmiddleware.SignMiddleware[Request](vmiddleware.NewSign().SetSecret(secret).
		SetSignKey("X-Sign").SetAlgorithm(vmiddleware.SignAlgorithmHmac).SetGetSignFunc(vmiddleware.GetSignByHeader)), func(c *vapi.Context) {
		voutput.O(c, nil, "test")
	})

	m := vconv.Map(Request{Name: "test"})
	signVal, err := vmiddleware.CalSign(m, secret, vmiddleware.SignAlgorithmHmac)

	vtest.Nil(t, err)
	vtest.Equal(t, "f62852695abd95e3703ce6b9903d2ce3616f078c317b02d0a67bde70c108afa9", signVal)

	dataStr, err := vjson.Encode(m)

	vtest.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(dataStr)))
	req.Header.Add("Content-Type", vapi.MimeJson)
	req.Header.Add("X-Sign", signVal)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
	vtest.Equal(t, "success", vjson.Parse(w.Body.String()).Get("msg").Str)
	vtest.Equal(t, "test", vjson.Parse(w.Body.String()).Get("data").Str)
}

func TestSignMiddlewareTraceIdMiddleware(t *testing.T) {
	traceId := "123456"
	g := vapi.Default()
	g.Any("/test", vmiddleware.TraceIdMiddleware[RequestTraceId](vmiddleware.DefaultTraceId), vmiddleware.SignMiddleware[RequestTraceId](vmiddleware.NewSign().SetSecret(secret)), func(c *vapi.Context) {
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
