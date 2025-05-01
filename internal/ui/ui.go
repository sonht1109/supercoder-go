package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/sonht1109/supercoder-go/internal/agent"
)

type TerminalChat struct{}

func clearScreen() {
	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func blue(text string) string {
	return "\033[34m" + text + "\033[0m"
}

func bold(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func underline(text string) string {
	return "\033[4m" + text + "\033[0m"
}

func printHeader(agent agent.BaseChatAgent) {
	clearScreen()
	fmt.Println(blue("█▀ █░█ █▀█ █▀▀ █▀█ █▀▀ █▀█ █▀▄ █▀▀ █▀█"))
	fmt.Println(blue("▄█ █▄█ █▀▀ ██▄ █▀▄ █▄▄ █▄█ █▄▀ ██▄ █▀▄"))
	fmt.Println(blue("v1.0.0")) // Replace with BuildInfo equivalent if needed
	fmt.Println()
	fmt.Println(blue("Model: " + agent.Model))
	fmt.Println(blue("Type '/help' for available commands.\n"))
}

func showHelp() {
	fmt.Println(underline("Available commands:"))
	fmt.Println("  " + bold("/help") + "  - Display this help message")
	fmt.Println("  " + bold("/clear") + " - Clear the terminal screen")
	fmt.Println("  " + bold("exit") + "\t- Terminate the chat session")
	fmt.Println("  " + bold("bye") + "\t- Terminate the chat session\n")
	fmt.Println("Just type any message to chat with the agent.")
	fmt.Println("To insert a new line in your message, use Shift+Enter (may vary by terminal).")
}

func Run(agent agent.BaseChatAgent) {
	printHeader(agent)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          bold("> "),
		HistoryFile:     "/tmp/readline.tmp", // Optional
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating readline:", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // Handle Ctrl+C or EOF
			fmt.Println(blue("\nChat session terminated. Goodbye!"))
			break
		}

		input := strings.TrimSpace(line)
		switch input {
		case "":
			continue
		case "/help":
			showHelp()
		case "/clear":
			clearScreen()
			printHeader(agent)
		case "exit", "bye":
			fmt.Println(blue("\nChat session terminated. Goodbye!"))
			return
		default:
			agent.Chat(input)
		}
	}
}
