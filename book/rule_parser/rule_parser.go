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

func (i Item) rules() (rules []rule.Rule) {

	if i.Warning != "" && i.Regex != "" {
		rules = append(rules, rule.Rule{i.Regex, i.Warning})
	} else if i.Group.Warning != "" {
		for _, regex := range i.Group.Regex {
			rules = append(rules, rule.Rule{regex, i.Group.Warning})
		}
	}

	return rules
}

type Item struct {
	Group struct {
		Warning string
		Regex   []string
	}
	Warning string
	Regex   string
}

func rulesInFile(filename string) (rules []rule.Rule) {
	m := struct{ Rules []Item }{}
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal([]byte(string(yamlFile)), &m)

	for _, item := range m.Rules {
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
