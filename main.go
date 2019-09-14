package main

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/ikeisuke/sample-github-tool/organization"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	o := organization.New(os.Getenv("GITHUB_ORGANIZATION"))
	o.SetGitHubV4Client(client)
	m := o.Members()
	m.Load()
}
