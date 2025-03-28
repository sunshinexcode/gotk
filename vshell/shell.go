package vshell

import "os/exec"

// Exec shell
func Exec(cmd string) (string, error) {
	f, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()

	return string(f), err
}
