package git

import (
	"os/exec"
	"strings"
)

type gitUC struct {
}

func NewGitUC() *gitUC {
	return &gitUC{}
}

func (uc *gitUC) FilesUnstaged() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(out))
	if trimmed == "" {
		return []string{}, nil
	}

	files := strings.Split(trimmed, "\n")
	return files, nil
}

func (uc *gitUC) GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	diff, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(diff), nil
}

func (uc *gitUC) Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (uc *gitUC) StageAllFiles() error {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}