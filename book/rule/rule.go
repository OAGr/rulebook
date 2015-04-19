package rule

import (
	"fmt"
	"regexp"
)

type Rule struct {
	Regex   string
	Warning string
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
