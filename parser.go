package main

import (
	//"fmt"
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

func ruleParser(dir string) []Rule {
	files, _ := ioutil.ReadDir(dir)

	var parseable_files []string

	for _, file := range files {
		if isParseable(file.Name()) {
			parseable_files = append(parseable_files, file.Name())
		}
	}

	rules := findFilesRules(parseable_files)
	return rules
}

func findFilesRules(filenames []string) []Rule {
	var _rules []Rule
	for _, filename := range filenames {
		_rules = append(_rules, findFileRules(filename)...)
	}
	return _rules
}

func findFileRules(filename string) []Rule {
	m := struct{ Rules []Item }{}
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal([]byte(string(yamlFile)), &m)

	var _rules []Rule
	for _, item := range m.Rules {
		_rules = append(_rules, parseItem(item)...)
	}

	return _rules
}

func isParseable(path string) bool {
	return (filepath.Ext(path) == ".yml")
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
