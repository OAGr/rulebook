package book

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Library struct {
	books []*Rulebook
}

func CurrentLibrary() (lib Library) {
	lib = Library{}
	lib.books = lib.findBooks()
	lib.setCurrentBook()
	return
}

func (l Library) String() string {
	decorated := make([]string, len(l.books))
	for i, s := range l.books {
		decorated[i] = s.decoratedName()
	}
	return strings.Join(decorated, "\n")
}

func (l Library) CurrentBook() *Rulebook {
	for _, b := range l.books {
		if b.IsCurrent {
			return b
		}
	}
	return &Rulebook{}
}

func (l Library) HasBook(book string) bool {
	return stringInSlice(book, l.bookNames())
}

func (l Library) GetBook(bookName string) *Rulebook {
	for _, s := range l.books {
		if s.Name == bookName {
			return s
		}
	}
	return &Rulebook{}
}

func LibraryPath() string {
	if os.Getenv("RULEBOOK_PATH") == "" {
		return os.Getenv("HOME") + "/.rulebooks/"
	} else {
		return os.Getenv("RULEBOOK_PATH")
	}
}

func (l Library) bookNames() (bookNames []string) {
	bookNames = make([]string, len(l.books))
	for i, s := range l.books {
		bookNames[i] = s.Name
	}
	return
}

// PreparationForLibrary
func (l *Library) setCurrentBook() {
	book := l.findCurrentBook()
	book.IsCurrent = true
	book.Rules = book.FindRules()
}

func (l Library) bookId(name string) (id int) {
	for i, b := range l.books {
		if b.Name == name {
			return i
		}
	}
	return 0
}

func (l Library) getRulebook(d DotRulebookFile) (book *Rulebook, err error) {
	file, _ := ioutil.ReadFile(d.path)
	rulebookName := strings.TrimSpace(string(file))
	if l.HasBook(rulebookName) {
		book = l.GetBook(rulebookName)
	}
	return book, err
}

func (l Library) findCurrentBook() *Rulebook {
	a := NewCurrentProject()
	b, err := l.getRulebook(a.DotRulebookFile)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func (l Library) findBooks() (rulebooks []*Rulebook) {
	bookPaths := getGitSubDirs(LibraryPath())
	rulebooks = make([]*Rulebook, len(bookPaths))

	for i, p := range bookPaths {
		name := BookPathToName(p, LibraryPath())
		rulebooks[i] = &Rulebook{Name: name}
	}
	return
}

func BookPathToName(path string, root string) string {
	return path[len(root) : len(path)-len(".git")-1]
}
