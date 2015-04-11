package main

import "regexp"

type Rule struct {
	Regex   string
	Warning string
}

func ViolatedLineRules(line string, rules []Rule) []Rule {
	var violated []Rule
	for _, rule := range rules {
		if DoesMatch(rule.Regex, line) {
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

func DoesMatch(reg string, text string) bool {
	regexp, _ := regexp.Compile(reg)
	match := regexp.FindString(text)
	if len(match) > 0 {
		return true
	} else {
		return false
	}
}
