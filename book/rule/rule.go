package rule

import (
	"fmt"
	"regexp"
)

type Rule struct {
	Regex   string
	Warning string
	Match   []string
	Nomatch []string
}

func (r Rule) String() string {
	return fmt.Sprintf("%s ->  %s", r.Regex, r.Warning)
}

func (r Rule) IsBrokenBy(s string) bool {
	return DoesMatch(r.Regex, s)
}

func DoesMatch(r string, s string) bool {
	regexp, _ := regexp.Compile(r)
	m := regexp.FindString(s)
	return len(m) > 0
}

func (r Rule) Test() []string {
	match, nomatch := r.failedTests()
	nomatch_0 := len(match)
	result := make([]string, (len(match) + len(nomatch)))

	for i, _ := range match {
		result[i] = "*match* did not match: Regex:" + r.Regex + " Phrase:" + match[i]
	}

	for i, _ := range nomatch {
		result[nomatch_0+i] = "*nomatch* did match: regex:" + r.Regex + " Phrase:" + nomatch[i]
	}

	return result
}

func (r Rule) failedTests() (match []string, nomatch []string) {
	for _, m := range r.Match {
		if !r.IsBrokenBy(m) {
			match = append(match, m)
		}
	}
	for _, n := range r.Nomatch {
		if r.IsBrokenBy(n) {
			nomatch = append(nomatch, n)
		}
	}
	return
}
