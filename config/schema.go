package config

type Config struct{
	Jira JiraConfig `json:"jira"`
}

type JiraConfig struct{
	Username string `json:"username"`
	Password string `json:"password"`
	BaseURL  string `json:"base_url"`
}