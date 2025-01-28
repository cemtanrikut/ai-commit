package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetAPIKey, gets OpenAI API key from os env
func GetAPIKey() string {
	// .env load
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set.")
		os.Exit(1)
	}

	return apiKey

}
