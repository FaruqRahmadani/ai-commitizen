package anthropic

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
)

type anthropicRepository struct {
	client anthropic.Client
}

func New(apiKey string) *anthropicRepository {
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &anthropicRepository{
		client: client,
	}
}

func (a *anthropicRepository) GenerateCommitMessage(input entity.CommitMessage) (string, error) {
	message, err := a.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(
				anthropic.NewTextBlock(fmt.Sprintf("Generate a concise, clear commit message that summarizes the changes in the provided git diff. Focus on the 'what' and 'why' of the changes. Return only the commit message text with no extra formatting, prefixes, or suffixes. Do not include ticket numbers, issue references, or conventional-commit type indicators. The response should be one line.\n\n%s", input.GitDiff)),
			),
		},
		Model: anthropic.ModelClaudeHaiku4_5,
	})

	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}
