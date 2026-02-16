package entity

import "github.com/faruqrahmadani/ai-commitizen/internal/constant"

type CommitMessage struct{
	TicketNumber string
	CommitType constant.CommitType
	GitDiff string
}