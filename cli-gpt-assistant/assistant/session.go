package assistant

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Manage chat history and reset/save

func RunAssistantLoop(apiKey, model string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			fmt.Println("Goodbye!")
			break
		}

		response, err := SendMessage(apiKey, model, input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Assistant:", response)
	}
	return nil
}
