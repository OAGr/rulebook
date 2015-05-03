package book

import (
	"github.com/oagr/rulebook/book/rule"
)

type EvaluateLine struct {
	line        string
	brokenRules []rule.Rule
	evaluator   *Evaluator
}

func (l *EvaluateLine) Evaluate() (err error) {
	for _, rule := range l.evaluator.Rulebook.Rules {
		isbroken, er := rule.IsBrokenBy(l.line)
		if er != nil {
			err = er
			return
		}
		if isbroken {
			l.brokenRules = append(l.brokenRules, rule)
		}
	}
	return
}
