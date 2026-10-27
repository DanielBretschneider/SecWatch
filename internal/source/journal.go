// secwatch/internal/source/journal.go
// Author: db
// Date: 2026-10-27
// Version: 0.1.5
// Description: Systemd journal and security event source integration

package source

import (
	"fmt"
	"os/exec"
)

// Return recent journal entries from current boot session
func ReadJournal(lines string) ([]byte, error) {
	// build command to fetch journal log lines
	cmd := exec.Command(
		"journalctl",
		"-b",
		"--no-pager",
		"-n",
		lines,
	)

	// cmd.Output() captures standard output as a byte slice
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("[-] read system journal: %w", err)
	}

	// return output (journald log line)
	return output, nil
}

// ReadFailedLogins returns journal entries matching common
// authentication failure messages from the current boot.
func ReadFailedLogins() ([]byte, error) {
	cmd := exec.Command(
		"journalctl",
		"-b",
		"--no-pager",
		"-g",
		"Failed password|authentication failure",
	)

	// get output from execute command
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("[-] reading failed login events failed: %w\n", err)
	}

	// return output
	return output, nil
}
