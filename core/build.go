package core

type Build struct {
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
	BuildDate  string
}

func NewBuild(commit, branch, state, summary, date string) *Build {
	b := new(Build)
	b.GitCommit = commit
	b.GitBranch = branch
	b.GitState = state
	b.GitSummary = summary
	b.BuildDate = date
	return b
}
