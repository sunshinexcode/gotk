package vapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Accounts    = gin.Accounts
	Context     = gin.Context
	Engine      = gin.Engine
	H           = gin.H
	HandlerFunc = gin.HandlerFunc
)

const (
	DebugMode   = gin.DebugMode
	ReleaseMode = gin.ReleaseMode
	TestMode    = gin.TestMode
)

func BasicAuth(accounts Accounts) HandlerFunc {
	return gin.BasicAuth(accounts)
}

func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	return gin.CreateTestContext(w)
}

func Default() *Engine {
	return gin.Default()
}

func SetMode(value string) {
	gin.SetMode(value)
}
