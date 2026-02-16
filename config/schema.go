package config

type Config struct{
	WithAI bool `json:"with_ai"`
	Jira JiraConfig `json:"jira"`
	Anthropic AnthropicConfig `json:"anthropic"`
}

type JiraConfig struct{
	Username string `json:"username"`
	Password string `json:"password"`
	BaseURL  string `json:"base_url"`
}

type AnthropicConfig struct{
	APIKey string `json:"api_key"`
}
