package github

import (
	"github.com/ikeisuke/sample-github-tool/github/organization"
	"github.com/ikeisuke/sample-github-tool/github/repository"
	"github.com/ikeisuke/sample-github-tool/github/upstream"
)

// GitHub root object
type GitHub struct {
	client upstream.GraphQL
}

func New() *GitHub {
	g := GitHub{
		client: &Client{},
	}
	return &g
}

func (g *GitHub) Organization(owner string) *organization.Organization {
	o := organization.New(owner)
	o.SetGitHubV4Client(g.client)
	return o
}

func (g *GitHub) Repository(owner string, name string) *repository.Repository {
	r := repository.New(owner, name)
	r.SetGitHubV4Client(g.client)
	return r
}
