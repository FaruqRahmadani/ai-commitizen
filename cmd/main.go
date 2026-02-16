package main

import (
	"fmt"
	"log"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/jira"
	"github.com/manifoldco/promptui"
)

type Service struct {
	JiraClient JiraUCItf
}

/*
	This apps will ack as a git commit message generator.
	It will prompt you to input the ticket number, then it will fetch the ticket summary from JIRA.
	After that, we will check your uncommitted changes.
	and we will generate a commit message with AI based on the changes.
	Then you should select the commitizen type such as: Feature, Fix, Chore, etc.
	Finally, it will generate the commit message like: <TICKET_NUMBER>: <COMMIT_TYPE> <COMMIT_MESSAGE>
	for Example: STOL-6969: (feat) Generate commit message with AI
*/
func main(){
	service := app()

	// ask your commit message
	ticketNumber := promptui.Prompt{
		Label:    "Ticket Number",
	}
	result, err := ticketNumber.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	ticket, err := service.JiraClient.GetTicket(result)
	if err != nil {
		log.Fatalf("failed to get ticket: %w", err)
	}

	fmt.Printf("You're working on %q\n", ticket.Summary)

	// check
}

func app() *Service{
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	jiraUC, err := jira.New(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.BaseURL)
	if err != nil {
		log.Fatalf("failed to init JIRA UC: %s", err)
	}

	return &Service{
		JiraClient: jiraUC,
	}
}