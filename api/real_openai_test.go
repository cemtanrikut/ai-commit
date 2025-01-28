package api

import (
	"fmt"
	"testing"
)

func TestGenerateCommitMessage_ModelNotFound(t *testing.T) {
	mockClient := &MockOpenAIClient{
		Err: fmt.Errorf("API error: The model `gpt-4` does not exist or you do not have access to it."),
	}

	diff := "diff --git a/main.go b/main.go\n+ Added a new feature"
	_, err := mockClient.GenerateCommitMessage(diff)
	if err == nil {
		t.Fatal("Expected error, got none")
	}

	expectedError := "API error: The model `gpt-4` does not exist or you do not have access to it."
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}
