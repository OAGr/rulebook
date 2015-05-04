package rule_parser

import (
	"github.com/oagr/rulebook/book/rule"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func RulesInDir(dir string) (rules []rule.Rule, err error) {
	fileList, err := findYml(dir)
	if err != nil {
		return
	}

	for _, filename := range fileList {
		file_rules, e := rulesInFile(filename)
		if e != nil {
			err = e
			break
		}
		rules = append(rules, file_rules...)
	}
	return
}

type Item struct {
	Group struct {
		Warning string
		Regex   []string
	}
	Warning string
	Regex   string
	Match   []string
	Nomatch []string
}

func (i Item) rules() (rules []rule.Rule) {

	if i.Warning != "" && i.Regex != "" {
		rule := rule.Rule{i.Regex, i.Warning, i.Match, i.Nomatch}
		rules = append(rules, rule)
	} else if i.Group.Warning != "" {
		for _, regex := range i.Group.Regex {
			rules = append(rules, rule.Rule{Regex: regex, Warning: i.Group.Warning})
		}
	}

	return rules
}

func rulesInFile(filename string) (rules []rule.Rule, err error) {
	m := struct{ Rules []Item }{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal([]byte(string(yamlFile)), &m)
	if err != nil {
		return
	}

	for _, item := range m.Rules {
		rules = append(rules, item.rules()...)
	}

	return
}

func findYml(dir string) (fileList []string, err error) {
	err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		matches, err := rule.DoesMatch(".yml", path)
		if err != nil {
			return err
		}
		if matches {
			fileList = append(fileList, path)
		}
		return nil
	})
	return
}
