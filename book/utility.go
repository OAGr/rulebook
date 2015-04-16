package book

import (
	"github.com/oagr/fs"
	"regexp"
)

func DoesMatch(r string, s string) bool {
	regexp, _ := regexp.Compile(r)
	m := regexp.FindString(s)
	return len(m) > 0
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
