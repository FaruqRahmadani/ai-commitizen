package jira

import (
	"log"

	jira "github.com/andygrunwald/go-jira"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
)

type jiraClient struct {
	Client *jira.Client
}

func New(username string, token string, baseURL string) *jiraClient {
	if username == "" || token == "" || baseURL == "" {
		return nil
	}

	jiraAuth := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}

	client, err := jira.NewClient(jiraAuth.Client(), baseURL)
	if err != nil {
		log.Fatalf("failed to init JIRA client: %s", err)
	}

	return &jiraClient{Client: client}
}

func (c *jiraClient) GetTicket(ticketNumber string) (*entity.JiraTicket, error) {
	issue, _, err := c.Client.Issue.Get(ticketNumber, nil)
	if err != nil {
		return nil, err
	}

	return &entity.JiraTicket{
		TicketType:   issue.Fields.Type.Name,
		TicketNumber: ticketNumber,
		Summary:      issue.Fields.Summary,
		Status:       issue.Fields.Status.Name,
	}, nil
}
