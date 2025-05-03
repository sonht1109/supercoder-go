package global

type Config struct {
	OpenAIAPIKey     string
	OpenAIAPIBaseURL string
	SearxngBaseURL   string
	Version          string
}

var Cfg Config
