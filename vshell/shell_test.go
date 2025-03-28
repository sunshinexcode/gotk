package vshell_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vshell"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestExec(t *testing.T) {
	res, err := vshell.Exec("pwd")

	vtest.Nil(t, err)
	vtest.Equal(t, true, res != "")
}
