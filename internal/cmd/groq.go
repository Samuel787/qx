package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GroqRequest struct {
	Model    string        `json:"model"`
	Messages []GroqMessage `json:"messages"`
}

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqResponse struct {
	Choices []GroqChoice `json:"choices"`
	Error   *GroqError   `json:"error,omitempty"`
}

type GroqChoice struct {
	Message GroqMessage `json:"message"`
}

type GroqError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func callGroqAPI(userQuery string) (string, error) {
	apiKey := os.Getenv("QX_GROQ_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("\033[33m⚠️  GROQ API Key Not Found\033[0m\n\nTo use qx, you need a GROQ API key. The good news: it's free with a generous tier!\n\n\033[1m1. Get your API key:\033[0m\n   Visit: https://console.groq.com\n   - Sign up (free, no credit card needed)\n   - Go to API Keys section\n   - Create a new key\n\n\033[1m2. Set the API key:\033[0m\n   qx set-key <your-api-key>\n\n\033[36m💡 Free Tier: 30 requests/min, unlimited usage\033[0m\n")
	}

	systemPrompt := `You are a helpful CLI command assistant. When a user asks for a command, respond with ONLY the command itself, without any explanation or markdown formatting. For example:
- If asked "how to list files", respond with: ls -la
- If asked "how to find files named test", respond with: find . -name test
- If asked "how to install npm packages", respond with: npm install

Always respond with just the raw command, nothing else.`

	req := GroqRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []GroqMessage{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: userQuery,
			},
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error making API call: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var groqResp GroqResponse
	err = json.Unmarshal(body, &groqResp)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	if groqResp.Error != nil {
		return "", fmt.Errorf("API error: %s", groqResp.Error.Message)
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return groqResp.Choices[0].Message.Content, nil
}
