package vcmd

import (
	"fmt"
	"os/exec"
)

// PortInUse check port whether used
func PortInUse(port int) (bool, error) {
	if port < 1 || port > 65535 {
		return false, fmt.Errorf("invalid port number: %d", port)
	}

	output, err := exec.Command("sh", "-c", fmt.Sprintf("lsof -i:%d", port)).CombinedOutput()
	return len(output) > 0, err
}
