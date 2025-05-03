package config

import (
	"os"

	"github.com/sonht1109/supercoder-go/internal/global"
)

func NewConfig() {
	cfg := global.Config{
		OpenAIAPIKey:     getEnv("OPENAI_API_KEY", ""),
		OpenAIAPIBaseURL: getEnv("OPENAI_API_BASE_URL", "https://api.openai.com/v1"),
		SearxngBaseURL:   getEnv("SEARXNG_BASE_URL", ""),
		Version:          getEnv("VERSION", "v0.0.1"),
	}

	global.Cfg = cfg
}

func getEnv(primary, fallback string) string {
	if val := os.Getenv(primary); val != "" {
		return val
	}
	return fallback
}
