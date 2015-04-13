package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Rulebook"
	app.Usage = "Test code against simple rules"
	app.Commands = []cli.Command{
		{
			Name:    "validate",
			Aliases: []string{"v"},
			Usage:   "(default) Analyze code sent in via pipe command",
			Action: func(c *cli.Context) {
				evaluateStdin()
			},
		},
		{
			Name:    "diff",
			Aliases: []string{"v"},
			Usage:   "validates `git diff` in current git repo",
			Action: func(c *cli.Context) {
				diff, err := exec.Command("sh", "-c", "git diff --color").Output()
				if err != nil {
					fmt.Printf("Terrible error %e", err)
				} else {
					lines := strings.Split(string(diff), "\n")
					evaluateText(lines)
				}
			},
		},
		{
			Name:    "run",
			Aliases: []string{"s"},
			Usage:   "run server on provided port number",
			Subcommands: []cli.Command{
				{
					Name:  "server",
					Usage: "initialize server (not yet implemented)",
					Action: func(c *cli.Context) {
						fmt.Println(CurrentLibrary().String())
					},
				},
				{
					Name:  "comment",
					Usage: "use downloaded books",
				},
			},
		},
		{
			Name:    "book",
			Aliases: []string{"b"},
			Usage:   "Manage Rulebooks",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list downloaded books",
					Action: func(c *cli.Context) {
						fmt.Println(CurrentLibrary().String())
					},
				},
				{
					Name:    "rules",
					Aliases: []string{"r"},
					Usage:   "List current rules",
					Action: func(c *cli.Context) {
						rules := getRules()
						println("Rulebook Rules")
						println("")
						for _, rule := range rules {
							println(rule.String())
						}
						println("Server running on port %s ", c.Args().First())
					},
				},
				{
					Name:  "use",
					Usage: "use downloaded books",
					Action: func(c *cli.Context) {
						book := c.Args().First()
						if stringInSlice(book, getRulebookNames()) {
							useBook(book)
							println("Change ENV RULEBOOK to book %s", book)
						} else {
							println("No book %s", book)
						}
					},
				},
				{
					Name:  "update",
					Usage: "update current book",
					Action: func(c *cli.Context) {
						fmt.Println("will implement")
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func bookLocation() string {
	//RULEBOOKSPATH
	//RULEBOOK
	return "/home/ozzie/rulebooks/github/oagr/rulebook1"
}

func getRules() []Rule {
	return ruleParser(bookLocation())
}

func evaluateStdin() {
	s := bufio.NewScanner(os.Stdin)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	evaluateText(lines)
}

func evaluateText(lines []string) {
	rules := getRules()

	var output []string
	output = append(output, decoratedMessage(lines, rules, getMessageType(lines))...)
	output = append(output, violationSummary(ViolatedLinesRules(lines, rules))...)

	for _, line := range output {
		fmt.Println(line)
	}
}

func getMessageType(lines []string) string {
	if isDiff(lines) {
		return "diff"
	} else {
		return "normal"
	}
}

func decoratedMessage(lines []string, rules []Rule, messageType string) []string {
	var decoratedLines []string
	for _, line := range lines {
		decoratedLines = append(decoratedLines, line)

		violations := ViolatedLineRules(line, rules)
		if len(violations) > 0 && shouldCheck(line, messageType) {
			decoratedLines = append(decoratedLines, violationError(violations))
		}
	}
	return decoratedLines
}

func shouldCheck(line string, messageType string) bool {
	if messageType == "diff" {
		return isCommitAddition(line)
	} else {
		return true
	}
}

func isCommitAddition(line string) bool {
	isAddition := "^\\+"
	withoutColor := line[5:]
	return DoesMatch(isAddition, line) || DoesMatch(isAddition, withoutColor)
}

func isDiff(text []string) bool {
	return DoesMatch("diff", text[0])
}

func additions(input string) (ouput string) {
	return input
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
