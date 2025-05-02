package tools

import (
	"encoding/json"
	"os"
)

type FileReadToolArguments struct {
	FilePath string `json:"filepath"`
}

type FileReadTool struct{}

func NewFileReadTool() *FileReadTool {
	return &FileReadTool{}
}

func (t *FileReadTool) Execute(arguments string) string {
	var args FileReadToolArguments
	err := json.Unmarshal([]byte(arguments), &args)
	if err != nil {
		return "Error: Invalid arguments"
	}

	fileContent, err := os.ReadFile(args.FilePath)
	if err != nil {
		return "Error: Unable to read file"
	}

	return string(fileContent)
}
