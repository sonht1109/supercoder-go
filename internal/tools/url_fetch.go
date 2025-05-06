package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type URLFetchTool struct{}

type URLFetchToolArguments struct {
	URL string `json:"url"`
}

func NewURLFetchTool() *URLFetchTool {
	return &URLFetchTool{}
}

func (t *URLFetchTool) Execute(arguments map[string]any) string {

	jsonData, err := json.Marshal(arguments)
	if err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	var args URLFetchToolArguments
	if err := json.Unmarshal(jsonData, &args); err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	fmt.Printf(utils.Green("🌐 Fetching URL: %s\n"), args.URL)

	resp, err := http.Get(args.URL)
	if err != nil {
		return fmt.Sprintf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Sprintf("Error: Received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	return string(body)
}
