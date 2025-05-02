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

func (t *URLFetchTool) Execute(arguments string) string {
	var args URLFetchToolArguments
	if err := json.Unmarshal([]byte(arguments), &args); err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	fmt.Printf(utils.Green("🌐 Fetching URL: %s"), args.URL)

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
