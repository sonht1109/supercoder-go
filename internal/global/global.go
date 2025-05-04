package global

type Config struct {
	OpenAIAPIKey     string
	OpenAIAPIBaseURL string
	SearxngBaseURL   string
	Version          string
	Env              string
	Debug            bool
	Model            string
}

var Cfg Config
