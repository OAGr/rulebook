package book

import (
	"fmt"
)

type PRLine struct {
	EvaluateLine EvaluateLine
	file         string
	SHA          string
	lineNumber   int
}

func (l PRLine) String() (result string) {
	result = l.EvaluateLine.line
	if len(l.EvaluateLine.brokenRules) != 0 {
		result += "\n" + l.BrokenRuleMessage()
	}
	return
}

func (l PRLine) BrokenRuleMessage() string {
	s := "- Rulebook Violation  "
	for _, rule := range l.EvaluateLine.brokenRules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	return s
}
