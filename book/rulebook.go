package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule_parser"
	"os"
)

type Rulebook struct {
	Name string
}

func CurrentBook() Rulebook {
	return Rulebook{os.Getenv("RULEBOOK")}
}

func (b Rulebook) Rules() []Rule {
	return rule_parser.rulesInDir(b.path())
}

func (b Rulebook) Use() {
	println("Run this command:")
	printThis := fmt.Sprintf("echo \"export RULEBOOK=%s\" | source /dev/stdin", b.Name)
	println(printThis)
}

func (b Rulebook) decoratedName() (decorated string) {
	decorated = "  " + b.Name
	if b == CurrentBook() {
		decorated = "*" + decorated[1:]
	}
	return
}

func (b Rulebook) path() (path string) {
	return (LibraryPath() + b.Name)
}
