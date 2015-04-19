package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/oagr/rulebook/book"
	"io/ioutil"
	"os"
	"os/exec"
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
				fmt.Println(evaluateStdin())
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
					fmt.Println(evaluateText(string(diff)))
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
						fmt.Println(book.CurrentLibrary().String())
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
						fmt.Println(book.CurrentLibrary().String())
					},
				},
				{
					Name:    "rules",
					Aliases: []string{"r"},
					Usage:   "List current rules",
					Action: func(c *cli.Context) {
						rules := book.CurrentBook().Rules
						println("Rulebook Rules")
						println("")
						for _, rule := range rules {
							println(rule.String())
						}
					},
				},
				{
					Name:  "use",
					Usage: "use downloaded books",
					Action: func(c *cli.Context) {
						bookName := c.Args().First()
						if book.CurrentLibrary().HasBook(bookName) {
							book.Rulebook{Name: bookName}.Use()
						} else {
							println("No book %s", bookName)
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

func evaluateStdin() string {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	return evaluateText(string(bytes))
}

func evaluateText(text string) string {
	return book.CurrentBook().EvaluateText(text)
}
