package repository

import (
	"os"

	"github.com/ikeisuke/sample-github-tool/github/repository/collaborators"
	"github.com/ikeisuke/sample-github-tool/github/upstream"
	"github.com/k0kubun/pp"
	"github.com/shurcooL/githubv4"
)

// Repository Repository
type Repository struct {
	client upstream.GraphQL
	id     string
	owner  string
	name   string
}

// New Users
func New(owner string, name string) *Repository {
	return &Repository{
		owner: owner,
		name:  name,
	}
}

// SetGitHubV4Client set to repository
func (r *Repository) SetGitHubV4Client(client upstream.GraphQL) {
	r.client = client
}

// Collaborators members
func (r *Repository) Collaborators() *collaborators.Collaborators {
	c := collaborators.New(r.owner, r.name)
	c.SetGitHubV4Client(r.client)
	return c
}

// ID return ID
func (r *Repository) ID() string {
	if r.id == "" {
		var query struct {
			Repository struct {
				ID githubv4.String
			} `graphql:"repository(owner: $owner, name: $name)"`
		}
		variables := map[string]interface{}{
			"owner": githubv4.String(r.owner),
			"name":  githubv4.String(r.name),
		}
		err := r.client.Query(&query, variables)
		if err != nil {
			pp.Printf("%+v\n", err)
			os.Exit(1)
		}
		r.id = string(query.Repository.ID)
	}
	return r.id
}

// CreateIssue create issue
func (r *Repository) CreateIssue(title string, body string) {
	var mutation struct {
		CreateIssue struct {
			ClientMutationID githubv4.String
			Issue            struct {
				DatabaseID githubv4.Int
			}
		} `graphql:"createIssue(input: $input)"`
	}
	input := githubv4.CreateIssueInput{
		RepositoryID: githubv4.String(r.ID()),
		Title:        githubv4.String(title),
		Body:         githubv4.NewString(githubv4.String(body)),
	}
	err := r.client.Mutate(&mutation, input, nil)
	if err != nil {
		pp.Printf("%+v\n", err)
		return
	}
	pp.Printf("%+\n", mutation)
}
