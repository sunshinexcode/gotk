package vapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Type aliases for gin package types
type (
	// Accounts represents HTTP Basic Authentication credentials
	Accounts = gin.Accounts
	// Context is the gin context wrapper
	Context = gin.Context
	// Engine is the gin web framework instance
	Engine = gin.Engine
	// H is a shortcut for map[string]interface{}
	H = gin.H
	// HandlerFunc defines the handler used by gin middleware
	HandlerFunc = gin.HandlerFunc
)

// Gin mode constants
const (
	// DebugMode indicates gin is running in debug mode
	DebugMode = gin.DebugMode
	// ReleaseMode indicates gin is running in release mode
	ReleaseMode = gin.ReleaseMode
	// TestMode indicates gin is running in test mode
	TestMode = gin.TestMode
)

// BasicAuth returns a Basic HTTP Authorization middleware
// It takes a map of user/password pairs and returns a handler function
// that checks for valid credentials before allowing access
//
// Example:
//
//	accounts := gin.Accounts{
//	    "admin": "password",
//	}
//	router.Use(BasicAuth(accounts))
func BasicAuth(accounts Accounts) HandlerFunc {
	return gin.BasicAuth(accounts)
}

// CreateTestContext creates a test context and engine for testing
// It takes a ResponseWriter and returns a context and engine instance
// This is useful for testing HTTP handlers without starting a server
//
// Example:
//
//	w := httptest.NewRecorder()
//	c, r := CreateTestContext(w)
//	handler(c)
//	assert.Equal(t, 200, w.Code)
func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	return gin.CreateTestContext(w)
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached
// This is the recommended way to start a new gin application
//
// Example:
//
//	r := Default()
//	r.GET("/ping", func(c *Context) {
//	    c.JSON(200, gin.H{"message": "pong"})
//	})
//	r.Run()
func Default() *Engine {
	return gin.Default()
}

// SetMode sets gin's mode (debug/release/test)
// This affects logging and error handling behavior
//
// Example:
//
//	SetMode(ReleaseMode) // Disable debug logging
//	SetMode(DebugMode)   // Enable debug logging
func SetMode(value string) {
	gin.SetMode(value)
}
