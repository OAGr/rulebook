package book

import (
	"github.com/oagr/rulebook/book/rule"
)

type Evaluator struct {
	Rulebook *Rulebook
	lines    []*EvaluateLine
}

func (e *Evaluator) Evaluate() {
	for _, l := range e.lines {
		l.Evaluate()
	}
}

func (e Evaluator) brokenRules() (brokenRules []rule.Rule) {
	for _, line := range e.lines {
		brokenRules = append(brokenRules, line.brokenRules...)
	}
	return
}
