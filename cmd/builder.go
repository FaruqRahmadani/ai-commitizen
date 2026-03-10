package main

import (
	"errors"
	"fmt"

	"github.com/faruqrahmadani/ai-commitizen/internal/constant"
	"github.com/faruqrahmadani/ai-commitizen/internal/entity"
	"github.com/fatih/color"
)

var (
	ErrUserCancelled = errors.New("user cancelled")
)

type CommitState struct {
	TicketNumber string
	Ticket       *entity.JiraTicket
	CommitType   constant.CommitType
	Diff         string
	Message      string
}

type CommitBuilder struct {
	s     *service
	state *CommitState
	err   error
}

func NewCommitBuilder(s *service) *CommitBuilder {
	return &CommitBuilder{
		s:     s,
		state: &CommitState{},
	}
}

func (b *CommitBuilder) CheckUnstagedFiles() *CommitBuilder {
	if b.err != nil {
		return b
	}

	files, err := b.s.gitUseCase.FilesUnstaged()
	if err != nil {
		b.err = fmt.Errorf("failed to check unstaged files: %w", err)
		return b
	}

	if len(files) > 0 {
		color.Yellow("You have %d unstaged changes. Please stage them with 'git add' first.", len(files))
		for _, file := range files {
			color.Red("  %s", file)
		}

		stageAll, err := PromptStageAllFiles()
		if err != nil {
			b.err = fmt.Errorf("failed to prompt stage all files: %w", err)
			return b
		}

		if !stageAll {
			color.Yellow("Please stage all files first")
			b.err = ErrUserCancelled
			return b
		}

		if err := b.s.gitUseCase.StageAllFiles(); err != nil {
			b.err = fmt.Errorf("failed to stage all files: %w", err)
			return b
		}
	}

	return b
}

func (b *CommitBuilder) RetrieveTicketInfo() *CommitBuilder {
	if b.err != nil {
		return b
	}

	if b.s.jiraUseCase != nil {
		ticketNumber, err := PromptTicketNumber()
		if err != nil {
			color.Red("Prompt failed %v", err)
			b.err = err
			return b
		}

		b.state.TicketNumber = ticketNumber

		ticket, err := b.s.jiraUseCase.GetTicket(ticketNumber)
		if err == nil && ticket != nil {
			cyan := color.New(color.Bold, color.FgCyan).SprintFunc()
			green := color.New(color.Bold, color.FgGreen).SprintFunc()
			color.White("You're working on [%s] %s (%s)", cyan(ticket.TicketType), cyan(ticket.Summary), green(ticket.Status))
			b.state.Ticket = ticket
		}
	}

	return b
}

func (b *CommitBuilder) SelectCommitType() *CommitBuilder {
	if b.err != nil {
		return b
	}

	commitType, err := PromptCommitType()
	if err != nil {
		color.Red("Prompt failed %v", err)
		b.err = err
		return b
	}

	b.state.CommitType = commitType
	return b
}

func (b *CommitBuilder) LoadChanges() *CommitBuilder {
	if b.err != nil {
		return b
	}

	diff, err := b.s.gitUseCase.GetDiff()
	if err != nil {
		b.err = fmt.Errorf("failed to get git diff: %w", err)
		return b
	}

	if len(diff) == 0 {
		color.Yellow("No staged changes found. Please stage your changes with 'git add' first.")
		b.err = ErrUserCancelled
		return b
	}

	b.state.Diff = diff
	return b
}

func (b *CommitBuilder) GenerateMessage() *CommitBuilder {
	if b.err != nil {
		return b
	}

	commitMessage, err := b.s.commitUseCase.GenerateCommitMessage(entity.CommitMessage{
		TicketNumber: b.state.TicketNumber,
		CommitType:   b.state.CommitType,
		GitDiff:      b.state.Diff,
	})
	if err != nil {
		b.err = fmt.Errorf("failed to generate commit message: %w", err)
		return b
	}

	green := color.New(color.Bold, color.FgGreen).SprintFunc()
	color.Yellow("Generated commit message:\n  %s", green(commitMessage))

	b.state.Message = commitMessage
	return b
}

func (b *CommitBuilder) ConfirmAndCommit() *CommitBuilder {
	if b.err != nil {
		return b
	}

	for {
		action, err := PromptCommit()
		if err != nil {
			color.Red("Prompt failed %v", err)
			b.err = err
			return b
		}

		if action == "No" {
			color.Yellow("Commit canceled.")
			b.err = ErrUserCancelled
			return b
		}

		if action == "Edit" {
			editedMsg, err := PromptEditCommitMessage(b.state.Message)
			if err != nil {
				b.err = fmt.Errorf("failed to edit commit message: %w", err)
				return b
			}
			b.state.Message = editedMsg

			// Show the edited message
			green := color.New(color.Bold, color.FgGreen).SprintFunc()
			color.Yellow("Updated commit message:\n  %s", green(b.state.Message))
			continue // Ask for confirmation again
		}

		// Action is "Yes"
		break
	}

	if err := b.s.gitUseCase.Commit(b.state.Message); err != nil {
		b.err = fmt.Errorf("failed to commit: %w", err)
		return b
	}

	color.Green("Commit successful.")
	return b
}

func (b *CommitBuilder) Build() error {
	if b.err == ErrUserCancelled {
		return nil
	}
	return b.err
}
