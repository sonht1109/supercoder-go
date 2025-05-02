package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sonht1109/supercoder-go/internal/global"
	"github.com/sonht1109/supercoder-go/internal/tools"
	"github.com/sonht1109/supercoder-go/internal/utils"
)

type ToolCallDescription struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ChatAgent struct {
	Client         *openai.Client
	Model          string
	Prompt         string
	ChatHistories  []openai.ChatCompletionMessage
	AvailableTools map[string]tools.Tool
}

const BasePrompt = `
  # Tool calling
  For each function call, return a json object with function name and arguments within <@TOOL></@TOOL> XML tags:

  <@TOOL>
  {"name": <function-name>, "arguments": "<json-encoded-string-of-the-arguments>"}
  </@TOOL>

  The arguments value is ALWAYS a JSON-encoded string, when there is no arguments, use empty string "".

  For example:
  <@TOOL>
  {"name": "file_read", "arguments": "{\"filePath\": \"example.txt\"}"}
  </@TOOL>

  <@TOOL>
  {"name": "project_structure", "arguments": ""}
  </@TOOL>

  The client will response with <@TOOL_RESULT>[content]</@TOOL_RESULT> XML tags to provide the result of the function call.
  Use it to continue the conversation with the user.

  # Safety
  Please refuse to answer any unsafe or unethical requests.
  Do not execute any command that could harm the system or access sensitive information.
  When you want to execute some potentially unsafe command, please ask for user confirmation first before generating the tool call instruction.

  # Agent Instructions
`

func NewChatAgent(prompt string) *ChatAgent {
	client := openai.NewClient(global.Cfg.OpenAIAPIKey)
	return &ChatAgent{
		Client:        client,
		Prompt:        prompt,
		ChatHistories: []openai.ChatCompletionMessage{},
	}
}

func (agent *ChatAgent) AddMessageIntoHistory(message string, role string) {
	agent.ChatHistories = append(agent.ChatHistories, openai.ChatCompletionMessage{
		Role:    role,
		Content: message,
	})
}

func (agent *ChatAgent) Chat(message string) {

	if message != "" {
		agent.AddMessageIntoHistory(message, openai.ChatMessageRoleUser)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: BasePrompt + agent.Prompt,
		},
	}
	messages = append(messages, agent.ChatHistories...)

	stream, err := agent.Client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    agent.Model,
		Messages: messages,
		Stream:   true,
	})
	if err != nil {
		fmt.Println("Stream error:", err)
		return
	}
	defer stream.Close()

	var buffer strings.Builder
	var currentToolContent strings.Builder

	isInToolTag := false
	toolTagStart := "<@TOOL>"
	toolTagEnd := "</@TOOL>"

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		content := resp.Choices[0].Delta.Content

		fmt.Print(utils.Red(content))

		buffer.WriteString(content)

		// handle tool tag logic
		if isInToolTag {

			currentToolContent.WriteString(content)
			toolContent := currentToolContent.String()

			// if receiving full tool call content, handle it
			if strings.Contains(toolContent, toolTagEnd) {
				rawDesc := strings.TrimSpace(strings.Split(toolContent, toolTagEnd)[0])

				var toolDesc ToolCallDescription
				err := json.Unmarshal([]byte(rawDesc), &toolDesc)
				if err != nil {
					fmt.Println("Error unmarshalling tool call description:", err)
				}

				agent.handleToolCall(toolDesc)

				currentToolContent.Reset()
				buffer.Reset()
				isInToolTag = false
			}

		} else {
			bufferContent := buffer.String()
			isInToolTag = strings.Contains(bufferContent, toolTagStart)

			// print out content
			// fmt.Print(utils.Blue(content))
		}
	}

	fmt.Print("\n")
}

func (agent *ChatAgent) handleToolCall(toolDesc ToolCallDescription) {
	selectedTool, ok := agent.AvailableTools[toolDesc.Name]
	if !ok {
		fmt.Println(utils.Red("Tool not found: " + toolDesc.Name))
		return
	}

	toolRes := selectedTool.Execute(toolDesc.Arguments)

	agent.AddMessageIntoHistory(fmt.Sprintf("Calling %s tool ...", toolDesc.Name), openai.ChatMessageRoleAssistant)
	agent.AddMessageIntoHistory(
		fmt.Sprintf("<@TOOL_RESULT>%s</@TOOL_RESULT>", toolRes),
		openai.ChatMessageRoleUser,
	)
}
