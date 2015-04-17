package rule_parser

import (
	"github.com/oagr/rulebook/book/rule"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func RulesInDir(dir string) (rules []rule.Rule) {
	fileList := findYml(dir)

	for _, filename := range fileList {
		rules = append(rules, rulesInFile(filename)...)
	}
	return
}

type item struct {
	group struct {
		warning string
		regex   []string
	}
	warning string
	regex   string
}

func (i item) rules() (rules []rule.Rule) {

	if i.warning != "" && i.regex != "" {
		rules = append(rules, rule.Rule{i.regex, i.warning})
	} else if i.group.warning != "" {
		for _, regex := range i.group.regex {
			rules = append(rules, rule.Rule{regex, i.group.warning})
		}
	}

	return rules
}

func rulesInFile(filename string) (rules []rule.Rule) {
	m := struct{ Elements []item }{}
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal([]byte(string(yamlFile)), &m)

	for _, item := range m.Elements {
		rules = append(rules, item.rules()...)
	}

	return
}

func findYml(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if rule.DoesMatch(".yml", path) {
			fileList = append(fileList, path)
		}

		return nil
	})
	return fileList
}
