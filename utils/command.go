package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// IsRunningFromGoRun returns true if the program is running via `go run`, false otherwise.
func IsRunningFromGoRun() bool {
	exePath, err := os.Executable()
	if err != nil {
		return false
	}
	// Typical go run paths contain "/go-build" or end with "/__debug_bin"
	if strings.Contains(exePath, "/go-build") || strings.HasSuffix(exePath, "/__debug_bin") {
		return true
	}
	return false
}
