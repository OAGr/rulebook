package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//import "strings"

func main() {
	rules := ruleParser("./test.yml")
	s := bufio.NewScanner(os.Stdin)
	var allViolatedRules []Rule

	for s.Scan() {
		text := s.Text()
		lines := strings.Split(text, "\n")

		for _, line := range lines {
			fmt.Println(line)
			violations := ViolatedRules(line, rules)
			allViolatedRules = append(allViolatedRules, violations...)
			if len(violations) != 0 {
				fmt.Println(ShowViolations(violations))
			}
		}
	}

	fmt.Println(ShowManyViolations(allViolatedRules))
	//fmt.Println(rules)
	//commentOnPr("https://github.com/OAGr/frequency_list/pull/1")
}

func additions(input string) (ouput string) {
	return input
}
