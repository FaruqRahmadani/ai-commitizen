package gemini

import (
	"context"
	"fmt"

	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type geminiRepository struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func New(apiKey string, modelName string) (*geminiRepository, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	if modelName == "" {
		modelName = "models/gemini-2.5-flash" // default model
	}

	model := client.GenerativeModel(modelName)

	return &geminiRepository{
		client: client,
		model:  model,
	}, nil
}

func (r *geminiRepository) GenerateCommitMessage(input entity.CommitMessage) (string, error) {
	ctx := context.Background()

	resp, err := r.model.GenerateContent(ctx, genai.Text(fmt.Sprintf("Based on the following git diff, generate a concise commit message that describes the changes. Return only the commit message text without any prefix, suffix, or formatting. Do not include ticket number or commit type in your response.\n\n%s", input.GitDiff)))
	if err != nil {
		return "", err
	}

	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}

		for _, part := range cand.Content.Parts {
			text, ok := part.(genai.Text)
			if !ok {
				continue
			}

			return string(text), nil
		}
	}

	return "", fmt.Errorf("no text content in Gemini response")
}
