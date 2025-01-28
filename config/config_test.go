package config

import (
	"os"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping test: OPENAI_API_KEY environment variable is not set")
	}

	apiKey = GetAPIKey()
	if apiKey == "" {
		t.Errorf("Expected a non-empty API key, but got empty string")
	}
}
