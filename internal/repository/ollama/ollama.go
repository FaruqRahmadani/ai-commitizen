package ollama

import (
	"fmt"
	"strings"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/rozoomcool/go-ollama-sdk"
)

type ollamaRepository struct {
	client *ollama.OllamaClient
	model  string
}

func New(cfg *config.Config) *ollamaRepository {
	client := ollama.NewClient(cfg.Ollama.BaseURL)
	return &ollamaRepository{
		client: client,
		model:  cfg.Ollama.Model,
	}
}

func (r *ollamaRepository) GenerateCommitMessage(input entity.CommitMessage) (string, error) {
	resp, err := r.client.Generate(r.model,
		fmt.Sprintf(
			"Generate a concise commit message (≤72 chars) summarizing the diff. "+
				"Output only the message text—no quotes, prefixes, types, or tickets.\n\n%s",
			input.GitDiff,
		),
	)
	if err != nil {
		return "", err
	}

	resp = strings.Trim(resp, " \n\"")

	return resp, nil
}
