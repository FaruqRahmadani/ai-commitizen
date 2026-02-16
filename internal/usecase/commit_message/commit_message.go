package commitmessage

import (
	"fmt"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/manifoldco/promptui"
)

type commitMessageUC struct {
	withAI bool
	repo   AIModelRepoItf
}

func NewCommitMessageUC(cfg *config.Config, aiRepo AIModelRepoItf) *commitMessageUC {
	return &commitMessageUC{
		withAI: cfg.WithAI,
		repo:   aiRepo,
	}
}

func (uc *commitMessageUC) GenerateCommitMessage(input entity.CommitMessage) (string, error) {
	if uc.withAI {
		return uc.repo.GenerateCommitMessage(input)
	}
	
	prompt := promptui.Prompt{
		Label:     "Input commit message",
	}

	message, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	if message == "" {
		return "", fmt.Errorf("commit message is empty")
	}

	return message, nil
}