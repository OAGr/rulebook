package main

import (
	"fmt"
	"github.com/oagr/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Library struct {
	list []Rulebook
}

type Rulebook struct {
	name string
}

func CurrentLibrary() (lib Library) {
	rulebooks := getRulebooks()
	return Library{rulebooks}
}

func (lib Library) String() string {
	decorated := make([]string, len(lib.list))
	for i, s := range lib.list {
		decorated[i] = s.decoratedName()
	}
	return strings.Join(decorated, "\n")
}

func currentBook() Rulebook {
	return Rulebook{os.Getenv("RULEBOOK")}
}

func useBook(book string) {
	fmt.Println(book)
	os.Setenv("RULEBOOK", book)
	exec.Command("sh", "-c", "EXPORT AA=bar")
}

func (book Rulebook) decoratedName() (decorated string) {
	decorated = "  " + book.name
	if book == currentBook() {
		decorated = "*" + decorated[1:]
	}
	return
}

func rulebookDir() string {
	if os.Getenv("RULEBOOK_PATH") == "" {
		return os.Getenv("HOME") + "/.rulebooks/"
	} else {
		return os.Getenv("RULEBOOK_PATH")
	}
}

func (book Rulebook) path() (path string) {
	return (rulebookDir() + book.name)
}

func getRulebookNames() []string {
	rulebooks := getRulebooks()
	rulebookNames := make([]string, len(rulebooks))
	for i, s := range rulebooks {
		rulebookNames[i] = s.name
	}
	return rulebookNames
}

func getRulebooks() []Rulebook {
	var rulebooks []Rulebook
	for _, gitDir := range getGitSubDirs(rulebookDir()) {
		rulebooks = append(rulebooks, Rulebook{getBookDir(rulebookDir(), gitDir)})
	}
	return rulebooks
}

func getGitSubDirs(root string) []string {
	walker := fs.Walk(root)
	var gitDirs []string

	for walker.Step() {
		if walker.Stat().Name() == ".git" {
			gitDirs = append(gitDirs, walker.Path())
			walker.SkipDir()
		}
	}
	return gitDirs
}

func getBookDir(root string, gitDir string) string {
	return gitDir[len(root) : len(gitDir)-len(".git")-1]
}

func parentsChildren(parents []string) []string {
	var children []string
	for _, parent := range parents {
		children = append(children, parentChildren(parent)...)
	}
	return children
}

func parentChildren(parent string) []string {
	var children []string
	_children, _ := ioutil.ReadDir(parent)

	fmt.Println(parent, "foobar")
	fmt.Println(_children)
	for _, child := range _children {
		children = append(children, child.Name())
	}
	return children
}

func listAll(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	return fileList
}
