package vauth

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Global variables for basic authentication keys
var (
	// BasicAuthorizationKey is the key used to store the Authorization header value
	// This is typically used in the format: "Basic base64(username:password)"
	BasicAuthorizationKey = "Authorization"
	// BasicAuthUserNameKey is the key used to store the username from basic auth
	// This is extracted from the Authorization header
	BasicAuthUserNameKey = "business"
)

// GetBasicAuthorization retrieves the Authorization header value from the context
// It supports both standard context and gin.Context
//
// Example:
//
//	auth := GetBasicAuthorization(ctx)
//	if auth != "" {
//	    // Process authorization header
//	}
func GetBasicAuthorization(ctx context.Context) string {
	if _, ok := ctx.(*gin.Context); ok {
		return GetGinContextBasicAuthorization(ctx.(*gin.Context))
	}

	if val, ok := ctx.Value(ContextKey(BasicAuthorizationKey)).(string); ok {
		return val
	}

	return ""
}

// GetBasicAuthUserName retrieves the username from basic authentication
// It supports both standard context and gin.Context
//
// Example:
//
//	username := GetBasicAuthUserName(ctx)
//	if username != "" {
//	    // Process username
//	}
func GetBasicAuthUserName(ctx context.Context) string {
	if _, ok := ctx.(*gin.Context); ok {
		return GetGinContextBasicAuthUserName(ctx.(*gin.Context))
	}

	if val, ok := ctx.Value(ContextKey(BasicAuthUserNameKey)).(string); ok {
		return val
	}

	return ""
}

// GetBasicAuthUserNameKey returns the current key used for storing the username
// This can be useful for debugging or when you need to know the current key name
func GetBasicAuthUserNameKey() string {
	return BasicAuthUserNameKey
}

// GetGinContextBasicAuthorization retrieves the Authorization header value from a gin.Context
// This is an internal helper function used by GetBasicAuthorization
func GetGinContextBasicAuthorization(c *gin.Context) string {
	if authorization, existed := c.Get(BasicAuthorizationKey); existed {
		return authorization.(string)
	}

	return ""
}

// GetGinContextBasicAuthUserName retrieves the username from basic authentication in a gin.Context
// This is an internal helper function used by GetBasicAuthUserName
func GetGinContextBasicAuthUserName(c *gin.Context) string {
	if basicAuthUserName, existed := c.Get(BasicAuthUserNameKey); existed {
		return basicAuthUserName.(string)
	}

	return ""
}

// SetBasicAuthorization stores the Authorization header value in the context
// Returns a new context with the authorization value set
//
// Example:
//
//	ctx = SetBasicAuthorization(ctx, "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
func SetBasicAuthorization(ctx context.Context, authorization string) context.Context {
	return context.WithValue(ctx, ContextKey(BasicAuthorizationKey), authorization)
}

// SetBasicAuthUserName stores the username from basic authentication in the context
// Returns a new context with the username value set
//
// Example:
//
//	ctx = SetBasicAuthUserName(ctx, "john_doe")
func SetBasicAuthUserName(ctx context.Context, basicAuthUserName string) context.Context {
	return context.WithValue(ctx, ContextKey(BasicAuthUserNameKey), basicAuthUserName)
}

// SetBasicAuthUserNameKey updates the global key used for storing the username
// This allows customization of the key name if needed
//
// Example:
//
//	SetBasicAuthUserNameKey("user")
func SetBasicAuthUserNameKey(basicAuthUserNameKey string) {
	BasicAuthUserNameKey = basicAuthUserNameKey
}
