package config

import (
	"os"

	"github.com/sonht1109/supercoder-go/internal/global"
)

var Version string

func NewConfig() {
	cfg := global.Config{
		OpenAIAPIKey:     getEnv("OPENAI_API_KEY", ""),
		OpenAIAPIBaseURL: getEnv("OPENAI_API_BASE_URL", "https://api.openai.com/v1"),
		SearxngBaseURL:   getEnv("SEARXNG_BASE_URL", "https://searx.be"),
		Debug:            getEnv("DEBUG", "false") == "true",
		Model:            getEnv("MODEL", "gpt-4.1-nano"),
		Version:          Version,
	}

	global.Cfg = cfg
}

func getEnv(primary, fallback string) string {
	if val := os.Getenv(primary); val != "" {
		return val
	}
	return fallback
}
