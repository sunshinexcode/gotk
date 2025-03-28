package vauth_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vauth"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestContextKey(t *testing.T) {
	vtest.Equal(t, "test", vauth.ContextKey("test").String())
}
