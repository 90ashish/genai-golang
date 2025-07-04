# CLI GPT Assistant

A terminal-based GPT assistant written in Go, with in-memory chat history, built-in commands, and pluggable model support.

---

## 📝 Features

- **In-memory chat history** for each session  
- **Built-in commands**  
  - `/reset` — clear the conversation (keeps only the system prompt)  
  - `/save`  — write the full session to `chat_history.txt`  
  - `/exit`  — quit the program  
- **Model flag** (`--model`) to choose any OpenAI chat model (e.g. `gpt-4o`)  
- **Loads** `OPENAI_API_KEY` from a `.env` file  

---

## ⚙️ Prerequisites

- Go 1.20 or higher installed  
- A valid OpenAI API key  

---

## 🚀 Installation Steps

1. Clone or download the repository to your local machine  
2. Create a file named `.env` in the project root  
   - Add your OpenAI API key inside it as:  
     `OPENAI_API_KEY=your_openai_api_key_here`  
3. Open the project in your preferred Go development environment  
4. Ensure dependencies are downloaded using your Go environment’s package management

---

## ▶️ Usage

- Launch the program through your development environment or using a compiled executable  
- Provide the model name using the `--model` flag (e.g., `gpt-4o`)  
- Interact via terminal: type messages or use special commands like `/save`, `/reset`, or `/exit`  

---

## 📖 Example Session

```
[GPT Assistant] Model: gpt-4o
Type your message below (type /exit to quit)

You: What is the capital of France?
Assistant: The capital of France is Paris.

You: /save
Chat history saved to chat_history.txt.

You: /reset
Conversation context reset.

You: What did I ask before?
Assistant: I’m not sure—this is a new conversation!

You: /exit
Goodbye!
```


---

## 🔧 Troubleshooting & Extending

- To change the initial system prompt, modify the `ResetHistory()` function in `assistant/commands.go`  
- To switch models, adjust the `--model` argument in your configuration  
- To manage token limits, add logic to summarize or truncate older messages  
- Feel free to contribute improvements or bug fixes via issues and pull requests  
