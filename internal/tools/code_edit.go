package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type CodeEditToolArguments struct {
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

type CodeEditTool struct{}

func NewCodeEditTool() *CodeEditTool {
	return &CodeEditTool{}
}

func (t *CodeEditTool) Execute(arguments map[string]any) string {

	jsonData, err := json.Marshal(arguments)
	if err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	var args CodeEditToolArguments
	if err := json.Unmarshal(jsonData, &args); err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	fmt.Printf(utils.Green("✏️ Editing file: %s\n"), args.FilePath)

	// Ensure the directory exists
	dir := filepath.Dir(args.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Sprintf("Error creating directory: %v", err)
	}

	// Write the content to the file
	if err := os.WriteFile(args.FilePath, []byte(args.Content), 0644); err != nil {
		return fmt.Sprintf("Error editing file: %v", err)
	}

	return fmt.Sprintf("Edited file %s successfully", args.FilePath)
}
