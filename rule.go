package main

//import "fmt"
import "regexp"

type Rule struct {
	regex   string
	warning string
}

func defaults() (rules []Rule) {
	r := []Rule{
		Rule{regex: "sub", warning: "I HATE SUBSPECIES"},
		Rule{regex: "woodlands", warning: "I prefer the forests"},
		Rule{regex: "far north", warning: "north?? Why north??"},
		Rule{regex: "1960s", warning: "the worst era..."},
		Rule{regex: "trapped", warning: "trapped? those poor birds!"},
	}
	return r
}

func ViolatedRules(line string) []Rule {
	var rules []Rule
	for _, rule := range defaults() {
		if DoesMatch(rule.regex, line) {
			rules = append(rules, rule)
		}
	}
	return rules
}

//func MatchLines(rule Rule, text string) []int {
//lines := strings.Split(text, "\n")
//var matchedLines []int
//for index, line := range lines {
//if Match(rule, line) {
//matchedLines = append(matchedLines, index)
//}
//}
//return matchedLines
//}

func DoesMatch(reg string, text string) bool {
	regexp, _ := regexp.Compile(reg)
	match := regexp.FindString(text)
	if len(match) > 0 {
		return true
	} else {
		return false
	}
}
