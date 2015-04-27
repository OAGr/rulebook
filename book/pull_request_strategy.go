package book

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
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
	strategy := PRStrategy{evaluator: evaluator}
	strategy.Prepare(url)
	strategy.LoadEvaluator()
	strategy.evaluator.Evaluate()
	strategy.Merge()
	fmt.Println(strategy.PRLines[0])
	//return strategy.Render()
	return err
}

func (p *PRStrategy) Prepare(url string) error {
	fmt.Println(p)
	user, repo, _ := parseUrl(url)
	fmt.Println(user, repo)
	client, err := getClient()
	if err != nil {
		return err
	}
	commits, _, _ := client.PullRequests.ListCommits("oagr", "frequency_list", 1, nil)
	fmt.Println(commits)
	files, _, _ := client.PullRequests.ListFiles(user, repo, 1, nil)
	fmt.Println(files)
	//p.SHA = *commits[len(commits)-1].SHA
	//fmt.Println(p)
	//files, _, _ := client.PullRequests.ListFiles(user, repo, 1, nil)
	//p.files = files
	//fmt.Println(p)
	return nil
}

//func (p *PRStrategy) LoadEvaluator() {
//for _, f := range p.files {
//lines := strings.Split(f.Patch, "\n")
//EvaluateLines := make(EvaluateLine, len(lines))
//PRLines := make(PRLine, len(lines))
//for i, l := range lines {
//EvaluateLines[i] = EvaluateLine{line: l}
//PrLines[i] = PrLine{SHA: p.SHA, file: f.Filename, lineNumber: i}
//}
//p.evaluator.lines = append(p.evaluator.lines, EvaluateLines)
//p.PRLines = append(p.PRLines, PRLines)
//}
//}

//func (p *PRStrategy) Merge() {
//for _, i := range p.PRLines {
//PrLines[i].EvaluateLine = EvaluateLine[i]
//}
//}
