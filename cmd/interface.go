package main

import "github.com/faruqrahmadani/ai-commitizen/internal/entity"

type JiraUCItf interface{
	GetTicket(ticketNumber string) (*entity.JiraTicket, error)
}