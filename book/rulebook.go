package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule"
	"github.com/oagr/rulebook/book/rule_parser"
	//"io/ioutil"
	"os"
	"os/exec"
	//"strings"
)

type Rulebook struct {
	Name      string
	Rules     []rule.Rule
	IsCurrent bool
}

func envBookName() string {
	return os.Getenv("RULEBOOK")
}

func (b *Rulebook) Change() {
	b.IsCurrent = true
}

func (b Rulebook) Update() {
	cmd, _ := exec.Command("git", "-C", b.path(), "pull", "origin", "master").Output()
	fmt.Println(string(cmd))
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
	if b.IsCurrent {
		decorated = "*" + decorated[1:]
	}
	return
}

func (b *Rulebook) makeCurrent() {
	b.IsCurrent = true
}

func (b Rulebook) path() (path string) {
	return (LibraryPath() + b.Name)
}
