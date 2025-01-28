package api

// GenerateCommitMessage, generates commit message with using OpenAI API
func GenerateCommitMessage(diff string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "system", "content": "Generate a commit message based on the following git diff."},
			{"role": "user", "content": diff},
		}
	}
}
