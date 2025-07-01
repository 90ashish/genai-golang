package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openAIKey := os.Getenv("OPENAI_API_KEY")
	if openAIKey == "" {
		log.Fatal("OPENAI_API_KEY not found in environment")
	}

	// Get input from CLI
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"your prompt here\"")
		return
	}
	userInput := strings.Join(os.Args[1:], " ")

	// Build request body
	reqBody := map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "user", "content": userInput},
		},
	}
	body, _ := json.Marshal(reqBody)

	// Prepare HTTP request
	req, _ := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+openAIKey)
	req.Header.Add("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}
	defer resp.Body.Close()

	// Read response body
	data, _ := io.ReadAll(resp.Body)

	// Parse and extract message content
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil || len(result.Choices) == 0 {
		log.Fatalf("Failed to parse response: %s", data)
	}

	// Print assistant's response
	fmt.Println("\n--- Response ---")
	fmt.Println(result.Choices[0].Message.Content)
}
