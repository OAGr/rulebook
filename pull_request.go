package main

import "github.com/google/go-github/github"
import "golang.org/x/oauth2"
import "fmt"
import "os"

//import "reflect"

func foobar() {
	client := getClient()

	commits, _, _ := client.PullRequests.ListCommits("oagr", "frequency_list", 1, nil)
	SHA := commits[len(commits)-1].SHA

	files, _, _ := client.PullRequests.ListFiles("oagr", "frequency_list", 1, nil)
	fmt.Println(files)
	var violations []Violation
	for _, file := range files {
		violations = append(violations, fileViolations(file)...)
	}

	for _, v := range violations {
		input := &github.PullRequestComment{
			Body:     github.String(v.rule.warning),
			CommitID: SHA,
			Path:     github.String(v.Filename),
			Position: github.Int(v.line),
		}
		comment, response, err := client.PullRequests.CreateComment("oagr", "frequency_list", 1, input)
		fmt.Println(comment, response, err)
	}

	fmt.Println(violations)
}

type tokenSource struct {
	token *oauth2.Token
}

// add Token() method to satisfy oauth2.TokenSource interface
func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func getClient() (client *github.Client) {
	code := os.Getenv("GITHUB_TEST_TOKEN")
	ts := &tokenSource{
		&oauth2.Token{AccessToken: code},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}
