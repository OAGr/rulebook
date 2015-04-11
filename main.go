package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sensei"
	app.Usage = "Test code for mistakes"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		println("Boom! I say")
	}
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "(default) Analyze code sent in via pipe command",
			Action: func(c *cli.Context) {
				evaluateStdin()
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "run server on provided port number",
			Action: func(c *cli.Context) {
				println("Server running on port %s ", c.Args().First())
			},
		},
		{
			Name:    "pull_request",
			Aliases: []string{"p"},
			Usage:   "Comment on a PR using local rules",
			Action: func(c *cli.Context) {
				println("Creating comments on PR", c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}

func getRules() []Rule {
	return ruleParser("./test.yml")
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
