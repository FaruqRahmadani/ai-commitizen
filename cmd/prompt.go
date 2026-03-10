package main

import (
	"os"
	"os/exec"
	"strings"

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

func PromptCommit() (string, error) {
	confirm := promptui.Select{
		Label: "Are you sure you want to commit with this message?",
		Items: []string{"Yes", "Edit", "No"},
	}
	_, resConfirm, err := confirm.Run()
	if err != nil {
		return "", err
	}

	return resConfirm, nil
}

func PromptStageAllFiles() (bool, error) {
	confirm := promptui.Select{
		Label: "Are you sure you want to stage all files?",
		Items: []string{"Yes", "No"},
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

func PromptEditCommitMessage(msg string) (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	tmpFile, err := os.CreateTemp("", "COMMIT_EDITMSG_*")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(msg); err != nil {
		return "", err
	}
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
