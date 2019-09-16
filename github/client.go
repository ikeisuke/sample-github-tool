package github

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var client = githubv4.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
)))

type Client struct {
}

func (c *Client) Query(q interface{}, variables map[string]interface{}) error {
	return client.Query(context.Background(), q, variables)
}

func (c *Client) Mutate(m interface{}, input githubv4.Input, variables map[string]interface{}) error {
	return client.Mutate(context.Background(), m, input, variables)
}
