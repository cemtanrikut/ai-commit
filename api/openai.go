package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"main.go/config"
)

// GenerateCommitMessage, generates commit message with using OpenAI API
func GenerateCommitMessage(diff string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// Explain to openai comment template
	systemMessage := `You are an AI assistant generating detailed Git commit messages. The messages should include:
	1. A short summary of the change (title).
	2. A longer, detailed explanation of what the change does (description).
	3. The context or reason for the change, if applicable.
	
	The commit message should follow this format:
	<type>: <short summary>
	
	<detailed description>
	
	Commit types:
	- feat: (new feature)
	- fix: (bug fix)
	- refactor: (refactoring production code)
	- style: (formatting, missing semi colons, etc; no code change)
	- docs: (changes to documentation)
	- test: (adding or refactoring tests; no production code change)
	- chore: (updating grunt tasks etc; no production code change)
	- wip: (work in progress; do not push!)`

	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "system", "content": systemMessage},
			{"role": "user", "content": diff},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to encode payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.GetAPIKey())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: ")
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return message, nil
}
