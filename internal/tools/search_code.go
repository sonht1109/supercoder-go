package tools

import (
	"encoding/json"
	"os/exec"
)

type SearchCodeTool struct{}

type SearchCodeToolArguments struct {
	Query string `json:"query"`
}

func NewSearchCodeTool() *SearchCodeTool {
	return &SearchCodeTool{}
}

func (s *SearchCodeTool) Execute(arguments string) string {
	var args SearchCodeToolArguments
	err := json.Unmarshal([]byte(arguments), &args)
	if err != nil {
		return "Error parsing arguments: " + err.Error()
	}

	cmd := exec.Command("git", "grep", "-n", "-i", args.Query, ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Error executing command: " + err.Error()
	}

	return string(output)
}
