package vfile

import (
	"github.com/gogf/gf/v2/os/gfile"
)

// Copy file/directory from <src> to <dst>.
//
// If <src> is file, it calls CopyFile to implements copy feature,
// or else it calls CopyDir.
func Copy(src string, dst string) error {
	return gfile.Copy(src, dst)
}

// Remove deletes all file/directory with `path` parameter.
// If parameter `path` is directory, it deletes it recursively.
//
// It does nothing if given `path` does not exist or is empty.
func Remove(path string) error {
	return gfile.Remove(path)
}

// ReplaceFile replaces content for file `path`.
func ReplaceFile(search, replace, path string) error {
	return gfile.ReplaceFile(search, replace, path)
}
