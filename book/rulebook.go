package book

import (
	"fmt"
	"github.com/oagr/rulebook/book/rule"
	"github.com/oagr/rulebook/book/rule_parser"
	//"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Rulebook struct {
	Name      string
	Rules     []rule.Rule
	IsCurrent bool
}

func (b *Rulebook) Change() {
	b.IsCurrent = true
}

func (b Rulebook) Update() {
	cmd, _ := exec.Command("git", "-C", b.path(), "pull", "origin", "master").Output()
	fmt.Println(strings.TrimSpace(string(cmd)))
}

func (b Rulebook) Clone() {
	name := strings.Replace(b.Name, "/", ":", 1)
	cloneFrom := "git@" + name + ".git"
	fmt.Println("Attempting to download", cloneFrom, "to library path", LibraryPath())
	_, err := exec.Command("git", "clone", cloneFrom, b.path()).Output()
	if err == nil {
		fmt.Println("Download was successful")
	} else {
		fmt.Println("Download failed with error", err)
	}
}

func (b *Rulebook) MakeCurrent() (err error) {
	b.IsCurrent = true
	if b.isDownloaded() {
		rules, err := b.FindRules()
		if err != nil {
			return err
		}
		b.Rules = rules
	}
	return err
}

func (b Rulebook) Use() {
	println("The rulebook can only be modified if not in a project with a .rulebook file.")
	println("To change rulebook, run this command:")
	printThis := fmt.Sprintf("export RULEBOOK=%s", b.Name)
	println(printThis)
}

func (b Rulebook) FindRules() ([]rule.Rule, error) {
	return rule_parser.RulesInDir(b.path())
}

func (b Rulebook) decoratedName() (decorated string) {
	decorated = "  " + b.Name
	if b.IsCurrent {
		decorated = "*" + decorated[1:]
	}
	if !b.isDownloaded() {
		decorated = decorated + " [Not Downloaded]"
	}
	return
}

func (b Rulebook) path() (path string) {
	return (LibraryPath() + b.Name)
}

func (b Rulebook) isDownloaded() bool {
	_, err := os.Stat(b.path())
	return (err == nil)
}

func (b Rulebook) Test() (string, error) {
	var result []string
	result = append(result, fmt.Sprintf("%-20s \t %-20s \t %-10s", "REGEX", "FAILURE", "TEST"))
	for _, rule := range b.Rules {
		test_result, err := rule.Test()
		if err != nil {
			return "", err
		}
		result = append(result, test_result...)
	}
	return strings.Join(result, "\n"), nil
}
