package book

import (
	"fmt"
	"github.com/google/go-github/github"
)

type PRLine struct {
	EvaluateLine EvaluateLine
	Filename     string
	SHA          string
	lineNumber   int
}

func (l PRLine) PRComment(SHA string) *github.PullRequestComment {
	return &github.PullRequestComment{
		Body:     github.String(l.String()),
		CommitID: github.String(SHA),
		Path:     github.String(l.Filename),
		Position: github.Int(l.lineNumber),
	}
}

func (l PRLine) String() (result string) {
	if len(l.EvaluateLine.brokenRules) != 0 {
		result += l.BrokenRuleMessage()
	}
	return
}

func (l PRLine) BrokenRuleMessage() string {
	s := "Rulebook Violations: "
	for _, rule := range l.EvaluateLine.brokenRules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	return s
}
