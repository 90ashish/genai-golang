package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GPT interaction logic

const openAIURL = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func SendMessage(apiKey, model, input string) (string, error) {
	messages := []Message{{Role: "user", Content: input}}
	reqBody := ChatRequest{Model: model, Messages: messages}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result ChatResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil || len(result.Choices) == 0 {
		return "", fmt.Errorf("invalid response: %s", string(respBody))
	}
	return result.Choices[0].Message.Content, nil
}
