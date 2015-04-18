package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule"
	"strings"
)

type EvaluateText struct {
	text        string
	Book        Rulebook
	messageType string
	lines       []EvaluateLine
}

func (t *EvaluateText) Evaluate() string {
	t.prepare()
	return t.String()
}

func (t EvaluateText) String() string {
	return t.body() + t.summary()
}

func (t *EvaluateText) prepare() {
	lines := strings.Split(t.text, "\n")
	elines := make([]EvaluateLine, len(lines))
	for i, line := range lines {
		eline := EvaluateLine{line: line, EvaluateText: *t}
		eline.Evaluate()
		elines[i] = eline
	}
}

func (t EvaluateText) body() string {
	decorated := make([]string, len(t.lines))
	for i, line := range t.lines {
		decorated[i] = line.String()
	}
	return strings.Join(decorated, "\n")
}

func (t EvaluateText) summary() string {
	rules := t.brokenRules()
	var s []string
	if len(rules) > 0 {
		s = append(s, "==================================================")
		s = append(s, fmt.Sprintf("%d Rulebook Violations ", len(rules)))
		s = append(s, "--------------------------------------------------")
		s = append(s, "\x1b[31;1m")
		for _, rule := range rules {
			regex := rule.Regex + "                     "
			regex = regex[0:20]
			s = append(s, fmt.Sprintf("regex: %s message: %s ", regex, rule.Warning))
		}
		s = append(s, "\x1b[0m")
		s = append(s, "==================================================")
	}
	return strings.Join(s, "\n")
}

func (t EvaluateText) brokenRules() (brokenRules []rule.Rule) {
	for _, line := range t.lines {
		brokenRules = append(brokenRules, line.brokenRules...)
	}
	return
}
