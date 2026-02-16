package main

import (
	"log"

	"github.com/faruqrahmadani/ai-commitizen/config"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/anthropic"
	"github.com/faruqrahmadani/ai-commitizen/internal/repository/gemini"
	commitmessage "github.com/faruqrahmadani/ai-commitizen/internal/usecase/commit_message"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/git"
	"github.com/faruqrahmadani/ai-commitizen/internal/usecase/jira"
	"github.com/fatih/color"
)

type Service struct {
	JiraUseCase   JiraUCItf
	GitUseCase    GitUCItf
	CommitUseCase CommitUCItf
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
func main() {
	service := app()

	cyan := color.New(color.Bold, color.FgCyan).SprintFunc()
	green := color.New(color.Bold, color.FgGreen).SprintFunc()

	// check if there are any unstaged changes
	files, err := service.GitUseCase.FilesUnstaged()
	if err != nil {
		log.Fatalf("failed to check unstaged files: %s", err)
	}

	if len(files) > 0 {
		color.Yellow("You have %d unstaged changes. Please stage them with 'git add' first.", len(files))
		for _, file := range files {
			color.Red("  %s", file)
		}

		stageAll, err := PromptStageAllFiles()
		if err != nil {
			log.Fatalf("failed to prompt stage all files: %s", err)
		}

		if !stageAll {
			color.Yellow("Please stage all files first")
			return
		}

		if err := service.GitUseCase.StageAllFiles(); err != nil {
			log.Fatalf("failed to stage all files: %s", err)
		}
	}

	// ask your commit message
	ticketNumber, err := PromptTicketNumber()
	if err != nil {
		color.Red("Prompt failed %v", err)
		return
	}

	ticket, err := service.JiraUseCase.GetTicket(ticketNumber)
	if err == nil && ticket != nil {
		color.White("You're working on [%s] %s (%s)", cyan(ticket.TicketType), cyan(ticket.Summary), green(ticket.Status))
	}

	// ask your commit type
	commitType, err := PromptCommitType()
	if err != nil {
		color.Red("Prompt failed %v", err)
		return
	}

	// check your uncommitted changes
	diff, err := service.GitUseCase.GetDiff()
	if err != nil {
		log.Fatalf("failed to get git diff: %v", err)
	}

	if len(diff) == 0 {
		color.Yellow("No staged changes found. Please stage your changes with 'git add' first.")
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

	color.Yellow("Generated commit message:\n  %s", green(commitMessage))

	// ask your confirmation
	confirm, err := PromptCommit()
	if err != nil {
		color.Red("Prompt failed %v", err)
		return
	}

	if !confirm {
		color.Yellow("Commit canceled.")
		return
	}

	// commit with the generated message
	if err := service.GitUseCase.Commit(commitMessage); err != nil {
		log.Fatalf("failed to commit: %s", err)
	}

	color.Green("Commit successful.")
}

func app() *Service {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	jiraUC, err := jira.New(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.BaseURL)
	if err != nil {
		log.Fatalf("failed to init JIRA UC: %s", err)
	}

	gitUC := git.NewGitUC()

	aiRepo := selectAIProvider(cfg)

	commitMessageUC := commitmessage.NewCommitMessageUC(cfg, aiRepo)

	return &Service{
		JiraUseCase:   jiraUC,
		GitUseCase:    gitUC,
		CommitUseCase: commitMessageUC,
	}
}

func selectAIProvider(cfg *config.Config) commitmessage.AIModelRepoItf {
	if !cfg.WithAI {
		return nil
	}

	switch cfg.Provider {
	case "", "Anthropic":
		return anthropic.New(cfg.Anthropic.APIKey)
	case "Gemini":
		repo, err := gemini.New(cfg.Gemini.APIKey, cfg.Gemini.Model)
		if err != nil {
			log.Fatalf("failed to init Gemini repository: %s", err)
		}

		return repo
	default:
		log.Fatalf("unsupported AI provider: %s", cfg.Provider)
		return nil
	}
}
