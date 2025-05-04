package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/sonht1109/supercoder-go/internal/agent"
	"github.com/sonht1109/supercoder-go/internal/global"
	"github.com/sonht1109/supercoder-go/internal/utils"
)

type TerminalChat struct{}

func clearScreen() {
	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func printHeader(agent agent.ChatAgent) {
	clearScreen()
	fmt.Println(utils.Blue("█▀ █░█ █▀█ █▀▀ █▀█ █▀▀ █▀█ █▀▄ █▀▀ █▀█"))
	fmt.Println(utils.Blue("▄█ █▄█ █▀▀ ██▄ █▀▄ █▄▄ █▄█ █▄▀ ██▄ █▀▄"))
	fmt.Println(utils.Blue(global.Cfg.Version)) // Replace with BuildInfo equivalent if needed
	fmt.Println()
	fmt.Println(utils.Blue("Model: " + agent.Model))
	fmt.Println(utils.Blue("Type '/help' for available commands.\n"))
}

func showHelp() {
	fmt.Println(utils.Underline("Available commands:"))
	fmt.Println("  " + utils.Bold("/help") + "  - Display this help message")
	fmt.Println("  " + utils.Bold("/clear") + " - Clear the terminal screen")
	fmt.Println("  " + utils.Bold("exit") + "\t- Terminate the chat session")
	fmt.Println("  " + utils.Bold("bye") + "\t- Terminate the chat session\n")
	fmt.Println("Just type any message to chat with the agent.")
	fmt.Println("To insert a new line in your message, use Shift+Enter (may vary by terminal).")
}

func Run(agent agent.ChatAgent) {
	printHeader(agent)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          utils.Bold("> "),
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
			fmt.Println(utils.Blue("\nChat session terminated. Goodbye!"))
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
			fmt.Println(utils.Blue("\nChat session terminated. Goodbye!"))
			return
		default:
			agent.ChatStream(input)
		}
	}
}
