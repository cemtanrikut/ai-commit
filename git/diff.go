package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// Runs git commit script
func RunGitCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// User can edit message with an editor
func EditCommitMessage(initalMessage string) (string, error) {
	// Generate temp file
	tempFile, err := os.CreateTemp("", "commit-message-*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}

	// Delete temp file after process complete
	defer os.Remove(tempFile.Name())

	// Write initial message to file
	if _, err := tempFile.WriteString(initalMessage); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Get user's editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		// If editor doesn't choose select VIM
		editor = "vim"
	}

	// Open editor
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to open editor: %v", err)
	}

	// Read edited message from file
	editedMessageBytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to read edited message: %v", err)
	}

	return strings.TrimSpace(string(editedMessageBytes)), nil
}
