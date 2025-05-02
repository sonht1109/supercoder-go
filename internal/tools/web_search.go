package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type WebSearchTool struct{}

type WebSearchToolArguments struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}

var baseURL = "http://localhost:8081"

func NewWebSearchTool() *WebSearchTool {
	return &WebSearchTool{}
}

func (t *WebSearchTool) Execute(arguments string) string {
	var args WebSearchToolArguments
	err := json.Unmarshal([]byte(arguments), &args)
	if err != nil {
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

	fmt.Println("====>", resStr.String())

	return resStr.String()
}
