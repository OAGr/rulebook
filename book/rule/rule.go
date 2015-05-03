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

func (r Rule) IsBrokenBy(s string) (bool, error) {
	return DoesMatch(r.Regex, s)
}

func DoesMatch(r string, s string) (matches bool, err error) {
	regexp, err := regexp.Compile(r)
	if err != nil {
		return
	}
	m := regexp.FindString(s)
	matches = len(m) > 0
	return
}

func (r Rule) Test() (result []string) {
	match, nomatch, err := r.failedTests()
	if err != nil {
		return
	}
	nomatch_0 := len(match)
	result = make([]string, (len(match) + len(nomatch)))

	for i, _ := range match {
		result[i] = fmt.Sprintf("%-20s \t %-20s \t %-10s", r.Regex, "*match* did not match", match[i])
	}

	for i, _ := range nomatch {
		result[nomatch_0+i] = fmt.Sprintf("%-20s \t %-20s \t %-10s", r.Regex, "*nomatch* did match", nomatch[i])
	}

	return result
}

func (r Rule) failedTests() (match []string, nomatch []string, err error) {
	for _, m := range r.Match {
		broken, er := r.IsBrokenBy(m)
		if er != nil {
			err = er
			return
		}
		if !broken {
			match = append(match, m)
		}
	}
	for _, n := range r.Nomatch {
		isbroken, er := r.IsBrokenBy(n)
		err = er
		if err != nil {
			return
		}
		if isbroken {
			nomatch = append(nomatch, n)
		}
	}
	return
}
