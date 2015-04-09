package main

//import "fmt"
import "github.com/google/go-github/github"

func fileViolations(f github.CommitFile) []Violation {
	return FindViolations(*f.Filename, *f.SHA, *f.Patch)
}
