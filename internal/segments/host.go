package segments

import (
	"fmt"
	"os"
)

type Host struct{}

func (h Host) ColoredOutput() string {
	return h.Output()
}

// Len return length of string without invisible characters counted
func (h Host) Len() int {
	return len(h.Output())
}

func (h Host) Output() string {
	sshClient := os.Getenv("SSH_CLIENT")

	if sshClient != "" {
		hostnameValue, _ := os.Hostname()
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}