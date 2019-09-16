package organization

import (
	"github.com/ikeisuke/sample-github-tool/github/organization/members"
	"github.com/ikeisuke/sample-github-tool/github/upstream"
)

// Organization organization
type Organization struct {
	client upstream.GraphQL
	login  string
}

// New Users
func New(login string) *Organization {
	return &Organization{
		login: login,
	}
}

// SetGitHubV4Client aaa
func (o *Organization) SetGitHubV4Client(client upstream.GraphQL) {
	o.client = client
}

// Members members
func (o *Organization) Members() *members.Members {
	m := members.New(o.login)
	m.SetGitHubV4Client(o.client)
	return m
}
