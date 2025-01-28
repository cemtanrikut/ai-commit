package api

import (
	"fmt"
	"testing"
)

func TestGenerateCommitMessage_Success(t *testing.T) {
	mockClient := &MockOpenAIClient{
		Response: "feat: Add mock test support",
	}

	diff := "diff --git a/main.go b/main.go\n+ Added a new feature"
	commitMessage, err := mockClient.GenerateCommitMessage(diff)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedMessage := "feat: Add mock test support"
	if commitMessage != expectedMessage {
		t.Errorf("Expected '%s', got '%s'", expectedMessage, commitMessage)
	}
}

func TestGenerateCommitMessage_Error(t *testing.T) {
	mockClient := &MockOpenAIClient{
		Err: fmt.Errorf("mock error"),
	}

	diff := "diff --git a/main.go b/main.go\n+ Added a new feature"
	_, err := mockClient.GenerateCommitMessage(diff)
	if err == nil {
		t.Fatal("Expected error, got none")
	}

	expectedError := "mock error"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}
