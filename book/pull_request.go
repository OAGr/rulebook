package book

import (
	"errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

//func commentOnPr(url string) {

////for _, v := range violations {
////input := &github.PullRequestComment{
////Body:     github.String(v.rule.Warning),
////CommitID: SHA,
////Path:     github.String(v.Filename),
////Position: github.Int(v.line),
////}
////comment, response, err := client.PullRequests.CreateComment(user, repo, 1, input)
////fmt.Println(comment, response, err)
////}
//}

func parseUrl(url string) (user string, repo string, pull_num string) {
	params := strings.Split(url[8:], "/")
	user = params[1]
	repo = params[2]
	pull_num = params[4]
	return user, repo, pull_num
}

type tokenSource struct {
	token *oauth2.Token
}

// add Token() method to satisfy oauth2.TokenSource interface
func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func getClient() (client *github.Client, err error) {
	code := os.Getenv("GITHUB_TEST_TOKEN")
	if code == "" {
		return nil, errors.New("Need GITHUB_TEST_TOKEN to use github client")
	}
	ts := &tokenSource{
		&oauth2.Token{AccessToken: code},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc), nil
}
