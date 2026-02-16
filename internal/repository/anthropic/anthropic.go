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

func New(apiKey string) *anthropicRepository{
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &anthropicRepository{
		client: client,
	}
}

func (a *anthropicRepository) GenerateCommitMessage(input entity.CommitMessage) (string, error){
	message, err := a.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(
				anthropic.NewTextBlock(fmt.Sprintf("Based on the following git diff, generate a concise commit message that describes the changes. Return only the commit message text without any prefix, suffix, or formatting. Do not include ticket number or commit type in your response.\n\n%s", input.GitDiff)),
			),
		},
		Model: anthropic.ModelClaudeHaiku4_5,
	})

	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}