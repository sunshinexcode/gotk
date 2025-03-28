package vdebug_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vdebug"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestD(t *testing.T) {
	vdebug.D("test", "test2")
}

func TestDebug(t *testing.T) {
	vdebug.Debug("test", "test2")
}

func TestDump(t *testing.T) {
	vdebug.Dump("test", "test2")
}

func TestStack(t *testing.T) {
	vtest.Equal(t, true, len(vdebug.Stack()) > 0)
}
