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
	app.Usage = "Simple Project Specific Linting"
	app.Authors = []cli.Author{cli.Author{Name: "Ozzie Gooen", Email: "ozzieagooen@gmail.com"}}
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:    "validate",
			Aliases: []string{"v"},
			Usage:   "Analyze content sent in via pipe command",
			Action: func(c *cli.Context) {
				str, err := evaluateStdin()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(str)
				}
			},
		},
		{
			Name:    "diff",
			Aliases: []string{"d"},
			Usage:   "Analyze content from `git diff` in current repository",
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
			Name:    "comment",
			Aliases: []string{"c"},
			Usage:   "Comment on a specific Github pull request",
			Action: func(c *cli.Context) {
				url := c.Args().First()
				library, err := book.CurrentLibrary()
				if err != nil {
					fmt.Println("PR Comment Failed:", err)
				}
				b := library.CurrentBook()
				err = book.ExecutePRStrategy(url, b)
				if err != nil {
					fmt.Println("PR Comment Failed:", err)
				}
			},
		},
		{
			Name:    "book",
			Aliases: []string{"b"},
			Usage:   "View & manage rulebooks",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all downloaded books",
					Action: func(c *cli.Context) {
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println("PR Comment Failed:", err)
						} else {
							fmt.Println(lib.String())
						}
					},
				},
				{
					Name:    "rules",
					Aliases: []string{"r"},
					Usage:   "List all rules in the current book",
					Action: func(c *cli.Context) {
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println("Rule List Fail:", err)
						} else {
							rules := lib.CurrentBook().Rules
							println("Rulebook Rules")
							for _, rule := range rules {
								println(rule.String())
							}
						}
					},
				},
				{
					Name:  "use",
					Usage: "Use a specific book",
					Action: func(c *cli.Context) {
						bookName := c.Args().First()
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println(err)
						} else {
							if lib.HasBook(bookName) {
								book.Rulebook{Name: bookName}.Use()
							} else {
								println("No book %s", bookName)
							}
						}
					},
				},
				{
					Name:  "clone",
					Usage: "Run `git clone` to get the currently selected book from Github",
					Action: func(c *cli.Context) {
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println(err)
						} else {
							lib.CurrentBook().Clone()
						}
					},
				},
				{
					Name:  "update",
					Usage: "Run `git pull` to update the currently selected book from Github",
					Action: func(c *cli.Context) {
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println(err)
						} else {
							lib.CurrentBook().Update()
						}
					},
				},
				{
					Name:    "test",
					Aliases: []string{"t"},
					Usage:   "Test Rulebook",
					Action: func(c *cli.Context) {
						lib, err := book.CurrentLibrary()
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(lib.CurrentBook().Test())
						}
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func evaluateStdin() (string, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return evaluateText(string(bytes))
}

func evaluateText(text string) (strategy string, err error) {
	lib, err := book.CurrentLibrary()
	if err != nil {
		fmt.Println(err)
	} else {
		b := lib.CurrentBook()
		strategy, err = book.ExecuteTextStrategy(text, b, "normal")
	}
	return
}
