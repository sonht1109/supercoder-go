package agent

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sonht1109/supercoder-go/internal/global"
)

type ToolCallDescription struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type BaseChatAgent struct {
	Client        *openai.Client
	Model         string
	Prompt        string
	ChatHistories []openai.ChatCompletionMessage
}

func NewBaseChatAgent(prompt string) *BaseChatAgent {
	client := openai.NewClient(global.Cfg.OpenAIAPIKey)
	return &BaseChatAgent{
		Client:        client,
		Model:         "gpt-4.1-nano",
		Prompt:        prompt,
		ChatHistories: []openai.ChatCompletionMessage{},
	}
}

func (agent *BaseChatAgent) AddUserMessage(message string) {
	agent.ChatHistories = append(agent.ChatHistories, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
}

func (agent *BaseChatAgent) Chat(message string) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: agent.Prompt,
		},
	}
	messages = append(messages, agent.ChatHistories...)

	fmt.Println("====> messages", messages)

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
	// var inToolTag bool
	// var toolTagEnd string

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		content := resp.Choices[0].Delta.Content
		buffer.WriteString(content)

		fmt.Print(content)

		// Handle tool tag logic
		// Simplified for example:
		// if strings.Contains(buffer.String(), "<@TOOL>") {
		// 	inToolTag = true
		// 	toolTagEnd = "</@TOOL>"
		// }
		// if inToolTag && strings.Contains(buffer.String(), toolTagEnd) {
		// 	inToolTag = false
		// 	fmt.Print("TOOL:", buffer.String())
		// 	buffer.Reset()
		// } else {
		// 	fmt.Print(content)
		// }
	}

	fmt.Print("\n")
	fmt.Println("====>", agent.ChatHistories)

	if message != "" {
		agent.AddUserMessage(message)
	}
}
