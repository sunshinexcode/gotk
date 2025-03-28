package vfx_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vfx"
)

func TestLifecycleMockAppend(t *testing.T) {
	l := &vfx.LifecycleMock{}

	l.Append(vfx.Hook{})
}
