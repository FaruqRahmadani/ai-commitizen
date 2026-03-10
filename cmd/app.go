package main

import (
	"log"
	"strings"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/anthropic"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/gemini"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/ollama"
	commitmessage "github.com/faruqrahmadani/ai-commitizen/internal/usecase/commit_message"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/git"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/jira"
)

func app() *service {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	jiraUC := jira.New(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.BaseURL)

	gitUC := git.NewGitUC()

	aiRepo := selectAIProvider(cfg)

	commitMessageUC := commitmessage.NewCommitMessageUC(cfg, aiRepo)

	return &service{
		jiraUseCase:   jiraUC,
		gitUseCase:    gitUC,
		commitUseCase: commitMessageUC,
	}
}

func selectAIProvider(cfg *config.Config) commitmessage.AIModelRepoItf {
	if !cfg.WithAI {
		return nil
	}

	switch strings.ToLower(cfg.Provider) {
	case "", "anthropic":
		return anthropic.New(cfg.Anthropic.APIKey)
	case "gemini":
		repo, err := gemini.New(cfg.Gemini.APIKey, cfg.Gemini.Model)
		if err != nil {
			log.Fatalf("failed to init Gemini repository: %s", err)
		}

		return repo

	case "ollama":
		return ollama.New(cfg)
	default:
		log.Fatalf("unsupported AI provider: %s", cfg.Provider)
		return nil
	}
}
