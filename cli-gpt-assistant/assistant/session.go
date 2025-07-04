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
	history := ResetHistory()

	for {
		fmt.Println("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Delegate to command handler
		h, handled, quit, err := HandleCommand(input, history)
		if handled {
			if err != nil {
				fmt.Println("Error:", err)
			}
			history = h
			if quit {
				fmt.Println("Goodbye!")
				return nil
			}
			continue
		}

		// Regular user input
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

		// Print and append assistant response
		fmt.Println("Assistant: ", response)
		history = append(history, Message{Role: "assistant", Content: response})
	}
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
