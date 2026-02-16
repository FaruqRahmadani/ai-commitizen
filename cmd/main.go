package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/constant"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/anthropic"
	commitmessage "github.com/faruqrahmadani/ai-commitizen/internal/usecase/commit_message"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/git"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/jira"
	"github.com/manifoldco/promptui"
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
	ticketNumber := promptui.Prompt{
		Label:    "Ticket Number",
	}
	ticketNumberStr, err := ticketNumber.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	ticket, err := service.JiraUseCase.GetTicket(ticketNumberStr)
	if err == nil && ticket != nil {
		fmt.Printf("You're working on %q\n", ticket.Summary)
	}

	// ask your commit type
	commitType := promptui.Select{
		Label: "Commit Type",
		Items: constant.CommitTypeItems,
	}
	commitTypeIndex, _, err := commitType.Run()
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

	// generate commit message with AI
	commitMessage, err := service.CommitUseCase.GenerateCommitMessage(entity.CommitMessage{
		TicketNumber: ticketNumberStr,
		CommitType:   constant.CommitTypeItems[commitTypeIndex],
		GitDiff:      diff,
	})
	if err != nil {
		log.Fatalf("failed to generate commit message: %s", err)
	}

	resultCommitMessage := fmt.Sprintf("%s: (%s) %s", ticketNumberStr, constant.CommitTypeItems[commitTypeIndex], commitMessage)
	
	fmt.Printf("%s\n", resultCommitMessage)

	// ask your confirmation
	confirm := promptui.Select{
		Label:     "Are you sure you want to commit with this message?",
		Items:     []string{"Yes", "No"},
	}
	_, resConfirm, err := confirm.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if resConfirm != "Yes" {
		fmt.Println("Commit canceled.")
		return
	}

	// commit with the generated message
	cmd := exec.Command("git", "commit", "-m", resultCommitMessage)
	err = cmd.Run()
	if err != nil {
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