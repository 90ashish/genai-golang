package assistant

import (
	"fmt"
	"os"
)

// Handle /reset, /save, /exit

// ResetHistory clears the conversation history, keeping only the system prompt.
func ResetHistory() []Message {
	return []Message{
		{Role: "system", Content: "You are a memory-enabled assistant. Keep track of the full conversation and answer follow-up questions based on history."},
	}
}

// SaveHistory writes the chat history to a file.
func SaveHistory(history []Message, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, msg := range history {
		if _, err := fmt.Fprintf(f, "%s: %s", msg.Role, msg.Content); err != nil {
			return err
		}
	}
	return nil
}

// HandleCommand processes built-in commands and returns updated history, whether it handled, if we should quit, and any error.
func HandleCommand(input string, history []Message) (newHistory []Message, handled bool, quit bool, err error) {
	newHistory = history
	switch input {
	case "/exit":
		// Quit CLI
		return newHistory, true, true, nil
	case "/reset":
		// Reset chat context
		newHistory = ResetHistory()
		fmt.Println("Conversation context reset.")
		return newHistory, true, false, nil
	case "/save":
		// Save chat history
		if err := SaveHistory(history, "chat_history.txt"); err != nil {
			return history, true, false, err
		}
		fmt.Println("Chat history saved to chat_history.txt.")
		return history, true, false, nil
	}
	return history, false, false, nil
}
