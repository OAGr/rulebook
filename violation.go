package main

//import "fmt"
import "strings"
import "fmt"

type Violation struct {
	rule     Rule
	Filename string
	SHA      string
	line     int
}

func ShowViolations(Rules []Rule) string {
	s := "\x1b[31;1m"
	s += "- VIOLATION ->  "
	for _, rule := range Rules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	s += "\x1b[0m"
	return s
}

func ShowManyViolations(Rules []Rule) string {
	s := ""
	s += fmt.Sprintf("==================================================\n")
	s += fmt.Sprintf("[Name Here]: %d Violations \n", len(Rules))
	s += fmt.Sprintf("--------------------------------------------------\n")
	s += "\x1b[31;1m"
	for _, rule := range Rules {
		regex := rule.Regex + "                "
		regex = regex[0:20]
		s += fmt.Sprintf("regex: %s message: %s \n", regex, rule.Warning)
	}
	s += "\x1b[0m"
	s += fmt.Sprintf("==================================================\n")
	return s
}

func FindViolations(Filename string, SHA string, Patch string, Rules []Rule) []Violation {
	var violations []Violation
	lines := strings.Split(Patch, "\n")
	for lineIndex, line := range lines {
		if !isCommitAddition(line) {
			continue
		}
		for _, rule := range ViolatedRules(line, Rules) {
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
