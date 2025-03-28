package vfile_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vfile"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestCopy(t *testing.T) {
	err := vfile.Copy("file.go", "/tmp/file.go")

	vtest.Nil(t, err)

	err = vfile.Copy("", "/tmp")

	vtest.NotNil(t, err)
	vtest.Equal(t, "source path cannot be empty", err.Error())
}

func TestRemove(t *testing.T) {
	vtest.Nil(t, vfile.Remove("/tmp/file_not_existed"))
}

func TestReplaceFile(t *testing.T) {
	vtest.NotNil(t, vfile.ReplaceFile("", "", ""))
}
