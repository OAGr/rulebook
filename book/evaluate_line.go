package book

import (
	"github.com/oagr/rulebook/book/rule"
)

type EvaluateLine struct {
	line        string
	brokenRules []rule.Rule
	evaluator   *Evaluator
}

func (l *EvaluateLine) Evaluate() {
	for _, rule := range l.evaluator.Rulebook.Rules {
		if rule.IsBrokenBy(l.line) {
			l.brokenRules = append(l.brokenRules, rule)
		}
	}
}
