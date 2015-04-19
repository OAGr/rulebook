package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule"
)

type EvaluateLine struct {
	line string
	EvaluateText
	brokenRules []rule.Rule
}

func (l *EvaluateLine) Evaluate() {
	for _, rule := range l.Book.Rules {
		if rule.IsBrokenBy(l.line) {
			l.brokenRules = append(l.brokenRules, rule)
		}
	}
}

func (l EvaluateLine) String() (result string) {
	result = l.line
	if l.shouldShowError() && len(l.brokenRules) != 0 {
		result += "\n" + l.BrokenRuleMessage()
	}
	return
}

func (l EvaluateLine) BrokenRuleMessage() string {
	s := "\x1b[31;1m"
	s += "- Rulebook Violation ->  "
	for _, rule := range l.brokenRules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	s += "\x1b[0m"
	return s
}

func (l EvaluateLine) shouldShowError() bool {
	return l.messageType != "diff" || l.isCommitAddition()
}

func (l EvaluateLine) isCommitAddition() bool {
	isAddition := "^\\+"
	withoutColor := l.line[5:]
	return DoesMatch(isAddition, l.line) || DoesMatch(isAddition, withoutColor)
}
