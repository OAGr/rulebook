package book

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"strings"
)

type PRStrategy struct {
	url       string
	SHA       string
	commits   []github.CommitFile
	client    *github.Client
	files     []github.CommitFile
	PRLines   []PRLine
	evaluator *Evaluator
}

func ExecutePRStrategy(url string, book *Rulebook) (err error) {
	if url == "" {
		return errors.New("input: Must give a url")
	} else if book == nil {
		return errors.New("input: Needs a book!")
	}
	evaluator := &Evaluator{Rulebook: book}
	strategy := PRStrategy{evaluator: evaluator, url: url}
	strategy.Prepare()
	strategy.LoadAndSplitEvaluator()
	strategy.evaluator.Evaluate()
	strategy.Merge()
	err = strategy.SendComments()
	return err
}

func (p *PRStrategy) SendComments() (err error) {
	user, repo, prNumber, err := parseUrl(p.url)
	if err != nil {
		return err
	}
	for _, l := range p.PRLines {
		if len(l.EvaluateLine.brokenRules) != 0 {
			input := l.PRComment(p.SHA)
			_, _, err = p.client.PullRequests.CreateComment(user, repo, prNumber, input)
			if err != nil {
				return err
			}
			fmt.Println("Posted Comment:", l.String())
		}
	}
	return
}

func (p *PRStrategy) Prepare() error {
	client, err := getClient()
	if err != nil {
		return err
	}
	user, repo, pull_num, err := parseUrl(p.url)
	if err != nil {
		return err
	}
	commits, _, _ := client.PullRequests.ListCommits(user, repo, pull_num, nil)
	if err != nil {
		return err
	}
	p.SHA = *commits[len(commits)-1].SHA
	p.client = client
	files, _, _ := client.PullRequests.ListFiles(user, repo, pull_num, nil)
	p.files = files
	return nil
}

func (p *PRStrategy) LoadAndSplitEvaluator() {
	for _, f := range p.files {
		lines := strings.Split(*f.Patch, "\n")
		EvaluateLines := make([]*EvaluateLine, len(lines))
		PRLines := make([]PRLine, len(lines))
		for i, l := range lines {
			EvaluateLines[i] = &EvaluateLine{line: l, evaluator: p.evaluator}
			PRLines[i] = PRLine{SHA: p.SHA, Filename: *f.Filename, lineNumber: i}
		}
		p.evaluator.lines = append(p.evaluator.lines, EvaluateLines...)
		p.PRLines = append(p.PRLines, PRLines...)
	}
}

func (p *PRStrategy) Merge() {
	for i, _ := range p.PRLines {
		p.PRLines[i].EvaluateLine = *p.evaluator.lines[i]
	}
}
