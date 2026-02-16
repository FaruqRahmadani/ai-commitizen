package jira

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
)

type jiraClient struct {
	Client *jira.Client
}

func New(username string, token string, baseURL string) (*jiraClient, error) {
	if username == "" || token == "" || baseURL == "" {
		return nil, fmt.Errorf("username, token, and baseURL are required")
	}

	jiraAuth := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}

	client, err := jira.NewClient(jiraAuth.Client(), baseURL)
	if err != nil {
		return nil, err
	}

	return &jiraClient{Client: client}, nil
}

func (c *jiraClient) GetTicket(ticketNumber string) (*entity.JiraTicket, error) {
	issue, _, err := c.Client.Issue.Get(ticketNumber, nil)
	if err != nil {
		return nil, err
	}

	return &entity.JiraTicket{
		TicketNumber: ticketNumber,
		Summary:      issue.Fields.Summary,
	}, nil
}