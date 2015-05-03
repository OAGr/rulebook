package book

import (
	"github.com/oagr/rulebook/book/rule"
)

type Evaluator struct {
	Rulebook *Rulebook
	lines    []*EvaluateLine
}

func (e *Evaluator) Evaluate() (err error) {
	for _, l := range e.lines {
		err = l.Evaluate()
	}
	return
}

func (e Evaluator) brokenRules() (brokenRules []rule.Rule) {
	for _, line := range e.lines {
		brokenRules = append(brokenRules, line.brokenRules...)
	}
	return
}
