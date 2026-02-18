package config

type Config struct {
	WithAI    bool
	Provider  string
	Jira      JiraConfig
	Anthropic AnthropicConfig
	Gemini    GeminiConfig
	Ollama    OllamaConfig
}

type JiraConfig struct {
	Username string
	Password string
	BaseURL  string
}

type AnthropicConfig struct {
	APIKey string
}

type GeminiConfig struct {
	APIKey string
	Model  string
}

type OllamaConfig struct {
	BaseURL string
	Model   string
}
