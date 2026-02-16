package main

import (
	"fmt"
	"log"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/anthropic"
	commitmessage "github.com/faruqrahmadani/ai-commitizen/internal/usecase/commit_message"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/git"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/jira"
)

type Service struct {
	JiraUseCase JiraUCItf
	GitUseCase  GitUCItf
	CommitUseCase   CommitUCItf
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
	ticketNumber, err := PromptTicketNumber()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	ticket, err := service.JiraUseCase.GetTicket(ticketNumber)
	if err == nil && ticket != nil {
		fmt.Printf("You're working on %q\n", ticket.Summary)
	}

	// ask your commit type
	commitType, err := PromptCommitType()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// check your uncommitted changes
	diff, err := service.GitUseCase.GetDiff()
	if err != nil {
		log.Fatalf("failed to get git diff: %v", err)
	}

	if len(diff) == 0 {
		fmt.Println("No staged changes found. Please stage your changes with 'git add' first.")
		return
	}

	// generate commit message
	commitMessage, err := service.CommitUseCase.GenerateCommitMessage(entity.CommitMessage{
		TicketNumber: ticketNumber,
		CommitType:   commitType,
		GitDiff:      diff,
	})
	if err != nil {
		log.Fatalf("failed to generate commit message: %s", err)
	}

	resultCommitMessage := fmt.Sprintf("%s: (%s) %s", ticketNumber, commitType, commitMessage)
	
	fmt.Printf("%s\n", resultCommitMessage)

	// ask your confirmation
	confirm, err := PromptCommit()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if !confirm {
		fmt.Println("Commit canceled.")
		return
	}

	// commit with the generated message
	if err := service.GitUseCase.Commit(resultCommitMessage); err != nil {
		log.Fatalf("failed to commit: %s", err)
	}

	fmt.Println("Commit successful.")
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

	gitUC := git.NewGitUC()

	anthropicRepo := anthropic.New(cfg.Anthropic.APIKey)

	commitMessageUC := commitmessage.NewCommitMessageUC(cfg, anthropicRepo)

	return &Service{
		JiraUseCase: jiraUC,
		GitUseCase:  gitUC,
		CommitUseCase:   commitMessageUC,
	}
}