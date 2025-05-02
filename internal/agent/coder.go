package agent

import (
	"fmt"

	"github.com/sonht1109/supercoder-go/internal/tools"
)

var basePrompt = fmt.Sprintf(`
  You are a senior software engineer AI agent. Your task is to help the user with their coding needs.
  You have access to the following tools:
  - %s: %s
  - %s: %s
  - %s: %s
  - %s: %s
  - %s: %s

  You can use these tools to help you with the user's request.

  When using the web-search tool, make sure you also use the url-fetch tool to read the content of the result URLs if needed.

  The discussion is about the code of the current project/folder. Always use the relevant tool to learn about the project if you are unsure before giving answer.
`, tools.CodeEditToolName, tools.CodeEditToolDescription,
	tools.FileReadToolName, tools.FileReadToolDescription,
	tools.CodeSearchToolName, tools.CodeSearchToolDescription,
	tools.WebSearchToolName, tools.WebSearchToolDescription,
	tools.URLFetchToolName, tools.URLFetchToolDescription,
)

type CoderAgent struct {
	ChatAgent
}

func NewCoderAgent(additionalPrompt string, model string) *CoderAgent {
	agent := &CoderAgent{
		ChatAgent: *NewChatAgent(basePrompt),
	}
	agent.Prompt = basePrompt + additionalPrompt
	agent.Model = model

	availableTools := make(map[string]tools.Tool)
	availableTools[tools.CodeEditToolName] = tools.NewCodeEditTool()
	availableTools[tools.FileReadToolName] = tools.NewFileReadTool()
	availableTools[tools.CodeSearchToolName] = tools.NewSearchCodeTool()
	availableTools[tools.WebSearchToolName] = tools.NewWebSearchTool()

	agent.AvailableTools = availableTools

	return agent
}
