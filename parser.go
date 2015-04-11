package main

import (
	"fmt"
	//"gopkg.in/yaml.v2"
	"github.com/kylelemons/go-gypsy/yaml"

	//"io/ioutil"
	"path/filepath"
	//"reflect"
)

type Rule struct {
	regex   string
	warning string
}

func pparse() {
	filename, _ := filepath.Abs("./test.yml")
	//yamlFile, err := ioutil.ReadFile(filename)
	yamlFile, _ := yaml.ReadFile(filename)

	fmt.Println(yamlFile.Count(""))
	//fmt.Println(yamlFile.Get("[1].message"))
	for i := 0; i < 4; i++ {
		str := fmt.Sprintf("[%d].message", 3)
		foo, _ := yamlFile.Get(str)
		fmt.Println(foo)
	}

	//mm := map(string:string)m.Root
	//typecast this into a map
	//for k, v := range Map(string: yaml.Map())m.Root {
	//fmt.Println(k, v)
	//}
	//fmt.Println(m)
	//for i, t := range m {

	//}

	//var rules []Rule
	//for i, t := range yamlFile {
	//fmt.Println(i, t)
	//}
	//return rules
	//for _, rule := range defaults() {
	//fmt.Println(yamlFile, err)

	//fmt.Printf("Value: %#v\n", config.Firewall_network_rules)
}
