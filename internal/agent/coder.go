package agent

// var coderAgentPrompt = `
// You are a senior software engineer AI agent. Your task is to help the user with their coding needs.

// You have access to the following tools:

// - CodeSearch: Search for code across the project.
// - ProjectStructure: Analyze the project structure.
// - FileRead: Read the contents of a file.
// - CodeEdit: Make edits to source code files.
// - CommandExecution: Run shell commands.
// - UrlFetch: Fetch and read contents from URLs.
// - WebSearch: Perform a web search.

// You can use these tools to help you with the user's request.

// When using the web-search tool, make sure you also use the url-fetch tool to read the content of the result URLs if needed.

// The discussion is about the code of the current project/folder. Always use the relevant tool to learn about the
// project if you are unsure before giving an answer.
// `

var basePrompt = "You are a helpful assistant."

type CoderAgent struct {
	BaseChatAgent
	AvailableTools []any
}

func NewCoderAgent(additionalPrompt string, model string) *CoderAgent {
	agent := &CoderAgent{
		BaseChatAgent: *NewBaseChatAgent(basePrompt),
	}
	agent.Prompt = basePrompt + additionalPrompt
	agent.Model = model
	// agent.availableTools = []Tool{
	// 	CodeSearchTool{},
	// 	ProjectStructureTool{},
	// 	FileReadTool{},
	// 	CodeEditTool{},
	// 	CommandExecutionTool{},
	// 	UrlFetchTool{},
	// 	WebSearchTool{},
	// }
	return agent
}

// func (a *CoderAgent) ToolDefinitionList() []FunctionDefinition {
// 	defs := []FunctionDefinition{}
// 	for _, tool := range a.availableTools {
// 		defs = append(defs, tool.Definition())
// 	}
// 	return defs
// }

// func (a *CoderAgent) ToolExecution(call ToolCallDescription) (string, error) {
// 	for _, tool := range a.availableTools {
// 		if tool.Definition().Name == call.Name {
// 			return tool.Execute(call.Arguments)
// 		}
// 	}
// 	return "", fmt.Errorf("Tool %s not found", call.Name)
// }
