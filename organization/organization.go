package organization

import (
	"github.com/ikeisuke/sample-github-tool/organization/members"
	"github.com/shurcooL/githubv4"
)

// Organization organization
type Organization struct {
	client *githubv4.Client
	login  string
}

// New Users
func New(login string) *Organization {
	return &Organization{
		login: login,
	}
}

// SetGitHubV4Client aaa
func (o *Organization) SetGitHubV4Client(client *githubv4.Client) {
	o.client = client
}

// Members members
func (o *Organization) Members() *members.Members {
	m := members.New(o.login)
	m.SetGitHubV4Client(o.client)
	return m
}
