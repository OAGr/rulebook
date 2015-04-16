package book

//import "fmt"
import "strings"
import "fmt"

type Violation struct {
	rule     Rule
	Filename string
	SHA      string
	line     int
}

func ViolationError(Rules []Rule) string {
	s := "\x1b[31;1m"
	s += "- Rulebook Violation ->  "
	for _, rule := range Rules {
		s += fmt.Sprintf("{regex: %s, message: %s}", rule.Regex, rule.Warning)
	}
	s += "\x1b[0m"
	return s
}

func ViolationSummary(Rules []Rule) []string {
	var s []string
	if len(Rules) > 0 {
		s = append(s, "==================================================")
		s = append(s, fmt.Sprintf("%d Rulebook Violations ", len(Rules)))
		s = append(s, "--------------------------------------------------")
		s = append(s, "\x1b[31;1m")
		for _, rule := range Rules {
			regex := rule.Regex + "                     "
			regex = regex[0:20]
			s = append(s, fmt.Sprintf("regex: %s message: %s ", regex, rule.Warning))
		}
		s = append(s, "\x1b[0m")
		s = append(s, "==================================================")
	}
	return s
}

func IsCommitAddition(line string) bool {
	isAddition := "^\\+"
	withoutColor := line[5:]
	return DoesMatch(isAddition, line) || DoesMatch(isAddition, withoutColor)
}

func FindViolations(Filename string, SHA string, Patch string, Rules []Rule) []Violation {
	var violations []Violation
	lines := strings.Split(Patch, "\n")
	for lineIndex, line := range lines {
		if !IsCommitAddition(line) {
			continue
		}
		for _, rule := range ViolatedLineRules(line, Rules) {
			v := Violation{rule: rule, Filename: Filename, SHA: SHA, line: lineIndex}
			violations = append(violations, v)
		}
	}
	return violations
}
