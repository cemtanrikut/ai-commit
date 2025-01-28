package git

import (
	"fmt"
	"os/exec"
)

// GetDiff, gets diff from staged changes
func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run git diff: %v", err)
	}

	return string(output), nil
}
