package assistant

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Manage chat history and reset/save

// RunAssistantWithMemory runs the interactive loop with chat history
func RunAssistantWithMemory(apiKey, model string) error {
	reader := bufio.NewReader(os.Stdin)

	history := []Message{
		{Role: "system", Content: "You are a memory-enabled assistant. Keep track of the full conversation and answer follow-up questions based on history."},
	}

	for {
		fmt.Println("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			fmt.Println("Goodbye!")
			break
		}

		history = append(history, Message{Role: "user", Content: input})

		// DEBUG: print history being sent
		// fmt.Println("DEBUG - Sending history:")
		// for _, msg := range history {
		// 	fmt.Printf("%s: %s\n", msg.Role, msg.Content)
		// }

		response, err := sendWithHistory(apiKey, model, history)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Assistant: ", response)
		history = append(history, Message{Role: "assistant", Content: response})
	}
	return nil
}

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
