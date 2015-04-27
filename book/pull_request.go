package book

import (
	"errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strconv"
	"strings"
)

func parseUrl(url string) (user string, repo string, pull_num int, err error) {
	if len(url) < 18 || url[0:18] != "https://github.com" {
		err = errors.New("Comment URL must start with 'https://github.com'")
		return
	}
	params := strings.Split(url[8:], "/")
	if len(params) < 5 {
		err = errors.New("ParseError: could not parse url")
		return
	}
	user = params[1]
	repo = params[2]
	pull_num, err = strconv.Atoi(params[4])
	if err != nil {
		return
	}
	return user, repo, pull_num, err
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
