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

func (l TextLine) String() (result string) {
	result = l.EvaluateLine.line
	if l.shouldShowError() && len(l.EvaluateLine.brokenRules) != 0 {
		result += "\n" + l.BrokenRuleMessage()
	}
	return
}
