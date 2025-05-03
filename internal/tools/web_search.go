package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sonht1109/supercoder-go/internal/global"
	"github.com/sonht1109/supercoder-go/internal/utils"
)

type WebSearchTool struct{}

type WebSearchToolArguments struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}

func NewWebSearchTool() *WebSearchTool {
	return &WebSearchTool{}
}

func (t *WebSearchTool) Execute(arguments map[string]any) string {

	jsonData, err := json.Marshal(arguments)
	if err != nil {
		return fmt.Sprintf("Error: Invalid arguments - %v", err)
	}

	baseURL := global.Cfg.SearxngBaseURL
	if baseURL == "" {
		return "Error: Searxng base URL is not set in the configuration."
	}

	var args WebSearchToolArguments
	if err := json.Unmarshal(jsonData, &args); err != nil {
		return "Error parsing arguments: " + err.Error()
	}

	fmt.Println()

	fmt.Println(utils.Green("Searching web with query: "), args.Query)

	params := url.Values{}
	params.Add("q", args.Query)
	params.Add("format", "json")

	searchURL := fmt.Sprintf("%s/search?%s", baseURL, params.Encode())

	resp, err := http.Get(searchURL)

	if err != nil {
		return "Error executing command: " + err.Error()
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "Error decoding response: " + err.Error()
	}

	var resStr strings.Builder

	if results, ok := result["results"].([]interface{}); ok {
		for _, r := range results {
			entry := r.(map[string]interface{})
			resStr.WriteString(fmt.Sprintf("Title: %s\nURL: %s\n\n", entry["title"], entry["url"]))
		}
	}

	return resStr.String()
}
