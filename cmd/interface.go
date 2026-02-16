package main

import "github.com/faruqrahmadani/ai-commitizen/internal/entity"

type JiraUCItf interface{
	GetTicket(ticketNumber string) (*entity.JiraTicket, error)
}

type GitUCItf interface{
	GetDiff() (string, error)
	Commit(msg string) error
}

type CommitUCItf interface{
	GenerateCommitMessage(input entity.CommitMessage) (string, error)
}