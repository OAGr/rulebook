package rule

import (
	"fmt"
)

type Rule struct {
	Regex   string
	Warning string
}

func (r Rule) String() string {
	return fmt.Sprintf("%s ->  %s", r.Regex, r.Warning)
}

func (r Rule) IsViolated(s string) bool {
	return DoesMatch(r.Regex, s)
}

func ViolatedLineRules(line string, rules []Rule) []Rule {
	var violated []Rule
	for _, rule := range rules {
		if rule.IsViolated(line) {
			violated = append(violated, rule)
		}
	}
	return violated
}

func ViolatedLinesRules(lines []string, rules []Rule) []Rule {
	var violated []Rule
	for _, line := range lines {
		violated = append(violated, ViolatedLineRules(line, rules)...)
	}
	return violated
}
