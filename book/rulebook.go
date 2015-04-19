package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule"
	"github.com/oagr/rulebook/book/rule_parser"
	"os"
)

type Rulebook struct {
	Name  string
	Rules []rule.Rule
}

func CurrentBook() Rulebook {
	rulebook := Rulebook{Name: currentBookName()}
	rulebook.Rules = rulebook.FindRules()
	return rulebook
}

func currentBookName() string {
	return os.Getenv("RULEBOOK")
}

func (b Rulebook) EvaluateText(text string) string {
	t := EvaluateText{text: text, Book: b}
	t.Evaluate()
	return t.String()
}

func (b Rulebook) Use() {
	println("Run this command:")
	printThis := fmt.Sprintf("echo \"export RULEBOOK=%s\" | source /dev/stdin", b.Name)
	println(printThis)
}

func (b Rulebook) FindRules() []rule.Rule {
	return rule_parser.RulesInDir(b.path())
}

func (b Rulebook) decoratedName() (decorated string) {
	decorated = "  " + b.Name
	if b.Name == currentBookName() {
		decorated = "*" + decorated[1:]
	}
	return
}

func (b Rulebook) path() (path string) {
	return (LibraryPath() + b.Name)
}
