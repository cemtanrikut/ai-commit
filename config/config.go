package config

import (
	"fmt"
	"os"
)

// GetAPIKey, gets OpenAI API key from os env
func GetAPIKey() string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set.")
		os.Exit(1)
	}

	return apiKey
}
