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
		return nil, fmt.Errorf("read system journal: %w", err)
	}

	// return output (journald log line)
	return output, nil
}
