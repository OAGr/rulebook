package book

import (
	"fmt"
	"strings"
)

type TextStrategy struct {
	text           string
	messageType    string
	formattedLines []TextLine
	evaluator      *Evaluator
}

func ExecuteTextStrategy(text string, book *Rulebook, messageType string) string {
	evaluator := &Evaluator{Rulebook: book}
	strategy := TextStrategy{text: text, messageType: messageType, evaluator: evaluator}
	strategy.Load()
	strategy.evaluator.Evaluate()
	strategy.Merge()
	return strategy.Render()
}

func (t *TextStrategy) Load() {
	lines := strings.Split(t.text, "\n")
	t.evaluator.lines = make([]*EvaluateLine, len(lines))
	for i, l := range lines {
		t.evaluator.lines[i] = &EvaluateLine{line: l, evaluator: t.evaluator}
	}
}

func (t *TextStrategy) Merge() {
	t.formattedLines = make([]TextLine, len(t.evaluator.lines))
	for i, l := range t.evaluator.lines {
		t.formattedLines[i].EvaluateLine = *l
	}
}

func (t TextStrategy) Render() string {
	return (t.body() + t.summary())
}

func (t TextStrategy) body() string {
	decorated := make([]string, len(t.formattedLines))
	for i, line := range t.formattedLines {
		decorated[i] = line.String()
	}
	return strings.Join(decorated, "\n")
}

func (t TextStrategy) summary() string {
	rules := t.evaluator.brokenRules()
	var s []string
	if len(rules) > 0 {
		s = append(s, "---------------------")
		s = append(s, fmt.Sprintf("\x1b[31;1m%d Rulebook Violations\x1b[0m", len(rules)))
		for _, rule := range rules {
			regex := rule.Regex + "                     "
			regex = regex[0:20]
			s = append(s, fmt.Sprintf("regex: %s message: %s ", regex, rule.Warning))
		}
		//s = append(s, "\x1b[0m")
	}
	validations := strings.Join(s, "\n")
	return validations
}
