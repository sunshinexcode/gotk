package vtest_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vtest"
)

func TestEqual(t *testing.T) {
	vtest.Equal(t, 1, 1)
}

func TestNil(t *testing.T) {
	vtest.Nil(t, nil)
}

func TestNotEqual(t *testing.T) {
	vtest.NotEqual(t, 1, 2)
}

func TestNotNil(t *testing.T) {
	vtest.NotNil(t, 1)
	vtest.NotNil(t, errors.New("test"))
}
