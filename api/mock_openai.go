package api

type MockOpenAIClient struct {
	Response string
	Err      error
}

func (m *MockOpenAIClient) GenerateCommitMessage(diff string) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Response, nil
}
