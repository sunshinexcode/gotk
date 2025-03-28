package vhttp_test

import (
	"strings"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestClient(t *testing.T) {
	resp, err := vhttp.C.R().EnableTrace().Get("https://httpbin.org/get")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())

	resp, err = vhttp.C.R().
		SetFormData(map[string]string{"name": "test"}).
		Post("http://httpbin.org/post")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())
	vtest.Equal(t, "test", vjson.Parse(resp.String()).Get("form.name").Str)

	resp, err = vhttp.C.R().
		ForceContentType("application/json").
		SetBody(`{"name":"test"}`).
		Post("http://httpbin.org/post")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())
	vtest.Equal(t, "test", vjson.Parse(resp.String()).Get("json.name").Str)
}

func TestNew(t *testing.T) {
	client := vhttp.NewClient()
	resp, err := client.R().EnableTrace().Get("https://httpbin.org/get")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())

	resp, err = client.SetTimeout(1 * time.Millisecond).R().Get("https://httpbin.org/get")

	vtest.NotNil(t, err)
	vtest.Equal(t, true, strings.Contains(err.Error(), "context deadline exceeded"))
	vtest.Equal(t, "1ms", client.GetClient().Timeout.String())
	vtest.Equal(t, "", string(resp.Body()))

	resp, err = client.R().Get("https://httpbin.org/get")

	vtest.NotNil(t, err)
	vtest.Equal(t, true, strings.Contains(err.Error(), "context deadline exceeded"))
	vtest.Equal(t, "1ms", client.GetClient().Timeout.String())
	vtest.Equal(t, "", string(resp.Body()))

	resp, err = client.SetTimeout(5 * time.Millisecond).R().Get("https://httpbin.org/get")

	vtest.NotNil(t, err)
	vtest.Equal(t, true, strings.Contains(err.Error(), "context deadline exceeded"))
	vtest.Equal(t, "5ms", client.GetClient().Timeout.String())
	vtest.Equal(t, "", string(resp.Body()))

	type Result struct {
		Data string
	}
	result := &Result{}
	resp, err = client.SetTimeout(5*time.Second).R().
		SetQueryParams(map[string]string{
			"limit": "20",
		}).
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name":"test"}`).
		SetResult(result).
		Post("http://httpbin.org/post")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())
	vtest.Equal(t, `{"name":"test"}`, result.Data)

	resp, err = client.SetTimeout(5 * time.Second).SetRetryCount(3).R().
		SetQueryParams(map[string]string{
			"limit": "20",
		}).
		ForceContentType("application/x-www-form-urlencoded").
		SetFormData(map[string]string{"name": "test"}).
		Post("http://httpbin.org/post")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())
	vtest.Equal(t, "test", vjson.Parse(resp.String()).Get("form.name").Str)

	resp, err = client.SetTimeout(5 * time.Second).R().
		SetQueryParams(map[string]string{
			"limit": "20",
		}).
		SetFormData(map[string]string{"name": "test"}).
		Post("http://httpbin.org/post")

	vtest.Nil(t, err)
	vtest.Equal(t, 200, resp.StatusCode())
	vtest.Equal(t, "200 OK", resp.Status())
	vtest.Equal(t, "test", vjson.Parse(resp.String()).Get("form.name").Str)
}
