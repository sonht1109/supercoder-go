package tools

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type SearchCodeTool struct{}

type SearchCodeToolArguments struct {
	Query string `json:"query"`
}

func NewSearchCodeTool() *SearchCodeTool {
	return &SearchCodeTool{}
}

func (s *SearchCodeTool) Execute(arguments map[string]any) string {

	jsonData, err := json.Marshal(arguments)
	if err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	var args SearchCodeToolArguments
	if err := json.Unmarshal(jsonData, &args); err != nil {
		return "Error parsing arguments: " + err.Error()
	}

	fmt.Printf(utils.Green("🔍 Search code for query: %s\n"), args.Query)

	cmd := exec.Command("git", "grep", "-n", "-i", args.Query, ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Error executing command: " + err.Error()
	}

	return string(output)
}
