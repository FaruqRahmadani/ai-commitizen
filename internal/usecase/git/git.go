package git

import "os/exec"

type gitUC struct {
}

func NewGitUC() *gitUC {
	return &gitUC{}
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