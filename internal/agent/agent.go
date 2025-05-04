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
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
	ID        string         `json:"id"`
}

type ChatAgent struct {
	Client         *openai.Client
	Model          string
	Prompt         string
	ChatHistories  []openai.ChatCompletionMessage
	AvailableTools map[string]tools.Tool
}

var basePrompt = `
  # Tool calling
  For each function call, you MUST return a json object with function name and arguments within <@TOOL></@TOOL> XML tags and follows format:

  <@TOOL>
  {"name": <function-name>, "arguments": <json-object>, "id": <function-id>}
  </@TOOL>

  The arguments value is ALWAYS a json object. When there is no arguments, use empty string "".

  For example:
  <@TOOL>
  {"name": "file_read", "arguments": {"filePath": "example.txt"}}
  </@TOOL>

  Function ID is a unique identifier for each function call. It is used to track the function call and its response. The ID should be a UUID v4 format.

  Do not hesitate to use the tools to help you with the user's request.

  # Safety
  Please refuse to answer any unsafe or unethical requests.
  Do not execute any command that could harm the system or access sensitive information.
  When you want to execute some potentially unsafe command, please ask for user confirmation first before generating the tool call instruction.

  Do not break any rules above, otherwise you will be fired.

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

func (agent *ChatAgent) AddMessageIntoHistory(
	message string,
	role string,
	toolCalls []openai.ToolCall,
	toolCallID string,
) {
	agent.ChatHistories = append(agent.ChatHistories, openai.ChatCompletionMessage{
		Role:       role,
		Content:    message,
		ToolCalls:  toolCalls,
		ToolCallID: toolCallID,
	})
}

func (agent *ChatAgent) ChatStream(message string) {

	if message != "" {
		agent.AddMessageIntoHistory(message, openai.ChatMessageRoleUser, nil, "")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: basePrompt + agent.Prompt,
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

	var checkToolContent strings.Builder
	var currentToolContent strings.Builder
	var tools []ToolCallDescription
	var currentContent strings.Builder
	var contentWithoutToolTag string

	isInToolTag := false
	toolTagStart := "<@TOOL>"
	toolTagEnd := "</@TOOL>"

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		content := resp.Choices[0].Delta.Content

		checkToolContent.WriteString(content)
		currentContent.WriteString(content)

		// handle tool tag logic
		if isInToolTag {

			currentToolContent.WriteString(content)
			toolContent := currentToolContent.String()

			// if receiving full tool call content, handle it
			if strings.Contains(toolContent, toolTagEnd) {
				if global.Cfg.Debug {
					fmt.Print(utils.Red(toolContent))
				}

				rawDesc := strings.TrimSpace(strings.Split(toolContent, toolTagEnd)[0])

				var toolDesc ToolCallDescription
				err := json.Unmarshal([]byte(rawDesc), &toolDesc)
				if err != nil {
					fmt.Println("Error unmarshalling tool call description:", err)
				}

				tools = append(tools, toolDesc)

				currentToolContent.Reset()
				checkToolContent.Reset()
				isInToolTag = false
			}

		} else {
			contentWithoutToolTag = contentWithoutToolTag + content
			bufferContent := checkToolContent.String()
			isInToolTag = strings.Contains(bufferContent, toolTagStart)

			if len(contentWithoutToolTag) > len(toolTagStart) {
				safeContent := contentWithoutToolTag[:len(contentWithoutToolTag)-len(toolTagStart)-1]

				// Print out content
				fmt.Print(utils.Blue(safeContent))
				contentWithoutToolTag = contentWithoutToolTag[len(safeContent):]
			}
		}
	}

	fmt.Println(utils.Blue(strings.Replace(contentWithoutToolTag, toolTagStart, "", -1)))

	agent.AddMessageIntoHistory(
		currentContent.String(),
		openai.ChatMessageRoleAssistant,
		nil,
		"",
	)

	for _, toolDesc := range tools {
		agent.handleToolCall(toolDesc)
	}
}

func (agent *ChatAgent) handleToolCall(toolDesc ToolCallDescription) {
	selectedTool, ok := agent.AvailableTools[toolDesc.Name]
	if !ok {
		fmt.Println(utils.Red("Tool not found: " + toolDesc.Name))
		return
	}

	toolRes := selectedTool.Execute(toolDesc.Arguments)

	encodedToolArgs, err := json.Marshal(toolDesc.Arguments)
	if err != nil {
		fmt.Println("Error marshalling tool arguments:", err)
		return
	}

	agent.AddMessageIntoHistory(
		fmt.Sprintf("Called %s tool with ID %s successfully", toolDesc.Name, toolDesc.ID),
		openai.ChatMessageRoleAssistant,
		[]openai.ToolCall{
			{ID: toolDesc.ID, Type: openai.ToolTypeFunction, Function: openai.FunctionCall{
				Name:      toolDesc.Name,
				Arguments: string(encodedToolArgs),
			}},
		},
		"",
	)

	agent.AddMessageIntoHistory(
		toolRes,
		openai.ChatMessageRoleTool,
		nil,
		toolDesc.ID,
	)

	fmt.Println(utils.Blue(toolRes))
}
