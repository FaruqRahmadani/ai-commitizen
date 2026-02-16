package constant

// CommitType represents the type of a commit message.
type CommitType string

const (
	CommitTypeFeature CommitType = "feat"
	CommitTypeFix     CommitType = "fix"
	CommitTypeChore   CommitType = "chore"
	CommitTypeDocs    CommitType = "docs"
	CommitTypeStyle   CommitType = "style"
	CommitTypeRefactor CommitType = "refactor"
	CommitTypeTest    CommitType = "test"
)

const (
	CommitTypeFeatureIndex = iota+1
	CommitTypeFixIndex
	CommitTypeChoreIndex
	CommitTypeDocsIndex
	CommitTypeStyleIndex
	CommitTypeRefactorIndex
	CommitTypeTestIndex
)

var CommitTypeItems = []CommitType{
	CommitTypeFeature,
	CommitTypeFix,
	CommitTypeChore,
	CommitTypeDocs,
	CommitTypeStyle,
	CommitTypeRefactor,
	CommitTypeTest,
}