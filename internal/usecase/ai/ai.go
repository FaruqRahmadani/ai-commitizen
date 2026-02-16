package ai

import (
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
)

type AIUC struct{
	repo AIModelRepoItf
}

func NewAIUC(repo AIModelRepoItf) *AIUC{
	return &AIUC{
		repo: repo,
	}
}

func (a *AIUC) GenerateCommitMessage(input entity.CommitMessage) (string, error){
	message, err := a.repo.GenerateCommitMessage(input)
	if err != nil {
		return "", err
	}

	return message, nil
}