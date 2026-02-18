package ollama

import (
	"fmt"

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
	resp, err := r.client.Generate(r.model, fmt.Sprintf("Based on the following git diff, generate a concise commit message that describes the changes. Return only the commit message text without any prefix, suffix, or formatting, and without any quotes on the first and last character. Do not include ticket number or commit type in your response.\n\n%s", input.GitDiff))
	if err != nil {
		return "", err
	}
		
	return resp, nil
}