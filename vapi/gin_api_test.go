package vapi_test

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestBasicAuth(t *testing.T) {
	vtest.Equal(t, "<gin.HandlerFunc Value>", reflect.ValueOf(vapi.BasicAuth(vapi.Accounts{"test": "pwd"})).String())
}

func TestCreateTestContext(t *testing.T) {
	c, e := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, c.Err())
	vtest.Equal(t, "/", e.BasePath())
}

func TestDefault(t *testing.T) {
	e := vapi.Default()

	vtest.Equal(t, "/", e.BasePath())
}

func TestSetMode(t *testing.T) {
	vapi.SetMode(vapi.DebugMode)
	vapi.SetMode(vapi.ReleaseMode)
	vapi.SetMode(vapi.TestMode)
}
