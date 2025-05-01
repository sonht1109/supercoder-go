package main

import (
	"github.com/sonht1109/supercoder-go/internal/agent"
	"github.com/sonht1109/supercoder-go/internal/config"
	"github.com/sonht1109/supercoder-go/internal/ui"
)

var AppConfig any

func main() {

	config.NewConfig()

	additionalPrompt := ""
	// if AppConfig.UseCursorRules {
	// 	additionalPrompt = LoadCursorRules()
	// }

	agent := agent.NewCoderAgent(additionalPrompt, "gpt-4o-mini")
	ui.Run(agent.BaseChatAgent)
}
