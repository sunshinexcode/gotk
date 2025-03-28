package vauth_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vauth"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestGetBasicAuthorization(t *testing.T) {
	ctx := context.TODO()

	vtest.Equal(t, "", vauth.GetBasicAuthorization(ctx))

	ctx = context.WithValue(ctx, vauth.ContextKey(vauth.BasicAuthorizationKey), "Basic test")

	vtest.Equal(t, "Basic test", vauth.GetBasicAuthorization(ctx))
}

func TestGetBasicAuthUserName(t *testing.T) {
	ctx := context.TODO()

	vtest.Equal(t, "", vauth.GetBasicAuthUserName(ctx))

	ctx = context.WithValue(ctx, vauth.ContextKey(vauth.BasicAuthUserNameKey), "test")

	vtest.Equal(t, "test", vauth.GetBasicAuthUserName(ctx))
}

func TestGetBasicAuthUserNameGinContext(t *testing.T) {
	ctx, _ := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, ctx.Err())
	vtest.Equal(t, "", vauth.GetBasicAuthUserName(ctx))

	ctx.Set(vauth.BasicAuthUserNameKey, "test")

	vtest.Equal(t, "test", vauth.GetBasicAuthUserName(ctx))
}

func TestGetBasicAuthUserNameKey(t *testing.T) {
	vtest.Equal(t, "business", vauth.GetBasicAuthUserNameKey())
}

func TestGetGinContextBasicAuthorization(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, c.Err())
	vtest.Equal(t, "", vauth.GetGinContextBasicAuthorization(c))

	c.Set(vauth.BasicAuthorizationKey, "Basic test")

	vtest.Equal(t, "Basic test", vauth.GetGinContextBasicAuthorization(c))
}

func TestGetGinContextBasicAuthUserName(t *testing.T) {
	c, _ := vapi.CreateTestContext(httptest.NewRecorder())

	vtest.Nil(t, c.Err())
	vtest.Equal(t, "", vauth.GetGinContextBasicAuthUserName(c))

	c.Set(vauth.BasicAuthUserNameKey, "test")

	vtest.Equal(t, "test", vauth.GetGinContextBasicAuthUserName(c))
}

func TestSetBasicAuthorization(t *testing.T) {
	ctx := vauth.SetBasicAuthorization(context.TODO(), "Basic test")

	vlog.Debug("TestSetBasicAuthorization", "vauth.GetBasicAuthorization(ctx)", vauth.GetBasicAuthorization(ctx))

	vtest.Equal(t, true, vauth.GetBasicAuthorization(ctx) != "")

	ctx = vauth.SetBasicAuthorization(context.TODO(), "Basic test")

	vtest.Equal(t, "Basic test", vauth.GetBasicAuthorization(ctx))
}

func TestSetBasicAuthUserName(t *testing.T) {
	ctx := vauth.SetBasicAuthUserName(context.TODO(), "test-")

	vlog.Debug("TestSetBasicAuthUserName", "vauth.GetBasicAuthUserName(ctx)", vauth.GetBasicAuthUserName(ctx))

	vtest.Equal(t, true, vauth.GetBasicAuthUserName(ctx) != "")

	ctx = vauth.SetBasicAuthUserName(context.TODO(), "test")

	vtest.Equal(t, "test", vauth.GetBasicAuthUserName(ctx))
}

func TestSetBasicAuthUserNameKey(t *testing.T) {
	vtest.Equal(t, "business", vauth.GetBasicAuthUserNameKey())

	vauth.SetBasicAuthUserNameKey("username")

	vtest.Equal(t, "username", vauth.GetBasicAuthUserNameKey())
}
