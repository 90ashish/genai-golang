package main

import (
	"cligptassistant/assistant"
	"cligptassistant/config"
	"flag"
	"fmt"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	apiKey := config.GetAPIKey()

	// Parse CLI flags
	model := flag.String("model", "gpt-4o", "OpenAI model to use (e.g., gpt-4, gpt-3.5-turbo)")
	flag.Parse()

	// Run the assistant loop with session context
	fmt.Printf("[GPT Assistant] Model: %s\nType your message below (type /exit to quit)\n\n", *model)
	err := assistant.RunAssistantWithMemory(apiKey, *model)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
