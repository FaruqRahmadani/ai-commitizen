package main

import (
	"github.com/faruqrahmadani/ai-commitizen/internal/constant"
	"github.com/manifoldco/promptui"
)

func PromptTicketNumber() (string, error) {
	prompt := promptui.Prompt{
		Label: "Ticket Number",
	}

	ticketNumber, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return ticketNumber, nil
}

func PromptCommitType() (constant.CommitType, error) {
	prompt := promptui.Select{
		Label: "Commit Type",
		Items: constant.CommitTypeItems,
	}

	_, commitType, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return constant.CommitType(commitType), nil
}

func PromptCommit() (bool, error) {
	confirm := promptui.Select{
		Label:     "Are you sure you want to commit with this message?",
		Items:     []string{"Yes", "No"},
	}
	_, resConfirm, err := confirm.Run()
	if err != nil {
		return false, err
	}

	if resConfirm != "Yes" {
		return false, nil
	}

	return true, nil
}

func PromptStageAllFiles() (bool, error) {
	confirm := promptui.Select{
		Label:     "Are you sure you want to stage all files?",
		Items:     []string{"Yes", "No"},
	}
	_, resConfirm, err := confirm.Run()
	if err != nil {
		return false, err
	}

	if resConfirm != "Yes" {
		return false, nil
	}

	return true, nil
}