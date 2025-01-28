package api

// OpenAIClient generates interface for talk with OpenAI API
type OpenAIClient interface {
	GenerateCommitMessage(diff string) (string, error)
}
