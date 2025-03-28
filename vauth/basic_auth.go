package vauth

import (
	"context"

	"github.com/gin-gonic/gin"
)

var (
	BasicAuthorizationKey = "Authorization"
	BasicAuthUserNameKey  = "business"
)

func GetBasicAuthorization(ctx context.Context) string {
	if _, ok := ctx.(*gin.Context); ok {
		return GetGinContextBasicAuthorization(ctx.(*gin.Context))
	}

	if val, ok := ctx.Value(ContextKey(BasicAuthorizationKey)).(string); ok {
		return val
	}

	return ""
}

func GetBasicAuthUserName(ctx context.Context) string {
	if _, ok := ctx.(*gin.Context); ok {
		return GetGinContextBasicAuthUserName(ctx.(*gin.Context))
	}

	if val, ok := ctx.Value(ContextKey(BasicAuthUserNameKey)).(string); ok {
		return val
	}

	return ""
}

func GetBasicAuthUserNameKey() string {
	return BasicAuthUserNameKey
}

func GetGinContextBasicAuthorization(c *gin.Context) string {
	if authorization, existed := c.Get(BasicAuthorizationKey); existed {
		return authorization.(string)
	}

	return ""
}

func GetGinContextBasicAuthUserName(c *gin.Context) string {
	if basicAuthUserName, existed := c.Get(BasicAuthUserNameKey); existed {
		return basicAuthUserName.(string)
	}

	return ""
}

func SetBasicAuthorization(ctx context.Context, authorization string) context.Context {
	return context.WithValue(ctx, ContextKey(BasicAuthorizationKey), authorization)
}

func SetBasicAuthUserName(ctx context.Context, basicAuthUserName string) context.Context {
	return context.WithValue(ctx, ContextKey(BasicAuthUserNameKey), basicAuthUserName)
}

func SetBasicAuthUserNameKey(basicAuthUserNameKey string) {
	BasicAuthUserNameKey = basicAuthUserNameKey
}
