package main

import (
	"log"
)

type service struct {
	jiraUseCase   JiraUCItf
	gitUseCase    GitUCItf
	commitUseCase CommitUCItf
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

	builder := NewCommitBuilder(service)
	err := builder.
		CheckUnstagedFiles().
		RetrieveTicketInfo().
		SelectCommitType().
		LoadChanges().
		GenerateMessage().
		ConfirmAndCommit().
		Build()

	if err != nil {
		log.Fatal(err)
	}
}
