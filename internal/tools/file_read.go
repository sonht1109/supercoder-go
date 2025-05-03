package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type FileReadToolArguments struct {
	FilePath string `json:"filePath"`
}

type FileReadTool struct{}

func NewFileReadTool() *FileReadTool {
	return &FileReadTool{}
}

func (t *FileReadTool) Execute(arguments map[string]any) string {

	jsonData, err := json.Marshal(arguments)
	if err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	var args FileReadToolArguments
	if err := json.Unmarshal(jsonData, &args); err != nil {
		return "Error: Invalid arguments"
	}

	fmt.Printf(utils.Green("📂 Reading file: %s\n"), args.FilePath)

	fileContent, err := os.ReadFile(args.FilePath)
	if err != nil {
		return "Error: Unable to read file"
	}

	return string(fileContent)
}
