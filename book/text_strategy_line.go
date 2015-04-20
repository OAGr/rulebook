package book

import (
	"fmt"
)

type TextLine struct {
	EvaluateLine EvaluateLine
	messageType  string
}

func (l TextLine) String() (result string) {
	result = l.EvaluateLine.line
	if l.shouldShowError() && len(l.EvaluateLine.brokenRules) != 0 {
		result += "\n" + l.BrokenRuleMessage()
	}
	return
}

func (l TextLine) BrokenRuleMessage() string {
	s := "\x1b[31;1m"
	s += "- Rulebook Violation ->  "
	for _, rule := range l.EvaluateLine.brokenRules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	s += "\x1b[0m"
	return s
}

func (l TextLine) shouldShowError() bool {
	return l.messageType != "diff" || l.isCommitAddition()
}

func (l TextLine) isCommitAddition() bool {
	isAddition := "^\\+"
	withoutColor := l.EvaluateLine.line[5:]
	return DoesMatch(isAddition, l.EvaluateLine.line) || DoesMatch(isAddition, withoutColor)
}
