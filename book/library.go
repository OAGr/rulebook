package book

import (
	"os"
	"strings"
)

type Library struct {
	books []Rulebook
}

func (l Library) String() string {
	decorated := make([]string, len(l.books))
	for i, s := range l.books {
		decorated[i] = s.decoratedName()
	}
	return strings.Join(decorated, "\n")
}

func (l Library) HasBook(book string) bool {
	bookNames := make([]string, len(l.books))
	for i, s := range l.books {
		bookNames[i] = s.Name
	}
	return stringInSlice(book, bookNames)
}

func CurrentLibrary() (lib Library) {
	return Library{findBooks()}
}

func LibraryPath() string {
	if os.Getenv("RULEBOOK_PATH") == "" {
		return os.Getenv("HOME") + "/.rulebooks/"
	} else {
		return os.Getenv("RULEBOOK_PATH")
	}
}

func findBooks() (rulebooks []Rulebook) {
	bookPaths := getGitSubDirs(LibraryPath())
	rulebooks = make([]Rulebook, len(bookPaths))

	for i, p := range bookPaths {
		name := BookPathToName(p, LibraryPath())
		rulebooks[i] = Rulebook{name}
	}
	return
}

func BookPathToName(path string, root string) string {
	return path[len(root) : len(path)-len(".git")-1]
}
