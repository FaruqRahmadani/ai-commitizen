package commitmessage

import (
	"fmt"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/constant"
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
		msg, err := uc.repo.GenerateCommitMessage(input)
		if err != nil {
			return "", err
		}

		return constructCommitMessage(input.TicketNumber, input.CommitType, msg), nil
	}

	prompt := promptui.Prompt{
		Label: "Input commit message",
	}

	message, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	if message == "" {
		return "", fmt.Errorf("commit message is empty")
	}

	return constructCommitMessage(input.TicketNumber, input.CommitType, message), nil
}

func constructCommitMessage(ticketNumber string, commitType constant.CommitType, commitMessage string) string {
	if ticketNumber == "" {
		return fmt.Sprintf("(%s) %s", commitType, commitMessage)
	}

	return fmt.Sprintf("%s: (%s) %s", ticketNumber, commitType, commitMessage)
}
