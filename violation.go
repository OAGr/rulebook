package main

//import "fmt"
import "strings"

type Violation struct {
	rule     Rule
	Filename string
	SHA      string
	line     int
}

func FindViolations(Filename string, SHA string, Patch string) []Violation {
	var violations []Violation
	lines := strings.Split(Patch, "\n")
	for lineIndex, line := range lines {
		if !isCommitAddition(line) {
			continue
		}
		for _, rule := range ViolatedRules(line) {
			v := Violation{rule: rule, Filename: Filename, SHA: SHA, line: lineIndex}
			violations = append(violations, v)
		}
	}
	return violations
}

func isCommitAddition(line string) bool {
	isAddition := "^\\+"
	return DoesMatch(isAddition, line)
}
