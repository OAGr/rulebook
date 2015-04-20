package book

import (
	"strings"
)

type PRStrategy struct {
	url       string
	SHA       string
	commits   []github.CommitFile
	files     *github.Client
	PRLines   []PRLine
	evaluator *Evaluator
}

func ExecutePRStrategy(url string, book *Rulebook) {
	evaluator := &Evaluator{Rulebook: book}
	strategy := PullRequestStrategy{text: text, messageType: messageType, evaluator: evaluator}
	strategy.Prepare()
	strategy.LoadEvaluator()
	strategy.evaluator.Evaluate()
	//strategy.Merge()
	//return strategy.Render()
}

func (p *PRStrategy) Prepare() {
	user, repo, _ := parseUrl(url)
	p.client = getClient()
	commits, _, _ := client.PullRequests.ListCommits(user, repo, 1, nil)
	p.SHA = commits[len(commits)-1].SHA
	files, _, _ := client.PullRequests.ListFiles(user, repo, 1, nil)
	p.files = files
}

func (p *PRStrategy) LoadEvaluator() {
	for _, f := range p.files {
		lines := strings.Split(f.Patch, "\n")
		EvaluateLines := make(EvaluateLine, len(lines))
		PRLines := make(PRLine, len(lines))
		for i, l := range lines {
			EvaluateLines[i] = EvaluateLine{line: l}
			PrLines[i] = PrLine{SHA: p.SHA, file: f.Filename, lineNumber: i}
		}
		p.evaluator.lines = append(p.evaluator.lines, EvaluateLines)
		p.PRLines = append(p.PRLines, PRLines)
	}
}

func (p *PRStrategy) Merge() {

}
