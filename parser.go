package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Item struct {
	Group struct {
		Warning string
		Regex   []string
	}
	Warning string
	Regex   string
}

func ruleParser(path string) (rules []Rule) {
	m := struct{ Rules []Item }{}
	filename, _ := filepath.Abs(path)
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal([]byte(string(yamlFile)), &m)

	var _rules []Rule
	for _, item := range m.Rules {
		_rules = append(_rules, parseItem(item)...)
	}

	return _rules
}

func parseItem(i Item) (r []Rule) {
	var rules []Rule

	if i.Warning != "" && i.Regex != "" {
		rules = append(rules, Rule{i.Regex, i.Warning})
	} else if i.Group.Warning != "" {
		for _, regex := range i.Group.Regex {
			rules = append(rules, Rule{regex, i.Group.Warning})
		}
	}

	return rules
}
