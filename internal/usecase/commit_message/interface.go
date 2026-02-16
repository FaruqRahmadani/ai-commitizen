package commitmessage

import "github.com/faruqrahmadani/ai-commitizen/internal/entity"

type AIModelRepoItf interface{
	GenerateCommitMessage(input entity.CommitMessage) (string, error)
}