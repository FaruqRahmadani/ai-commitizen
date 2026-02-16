package config

type Config struct{
	WithAI bool
	Jira JiraConfig
	Anthropic AnthropicConfig
}

type JiraConfig struct{
	Username string
	Password string
	BaseURL  string
}

type AnthropicConfig struct{
	APIKey string
}
