package collaborators

import (
	"os"

	"github.com/ikeisuke/sample-github-tool/github/upstream"
	"github.com/k0kubun/pp"
	"github.com/shurcooL/githubv4"
)

// Collaborators struct
type Collaborators struct {
	client        upstream.GraphQL
	owner         string
	name          string
	collaborators []Collaborator
}

type rateLimit struct {
	Cost      githubv4.Int
	Limit     githubv4.Int
	NodeCount githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}
type pageInfo struct {
	StartCursor     githubv4.String
	EndCursor       githubv4.String
	HasPreviousPage githubv4.Boolean
	HasNextPage     githubv4.Boolean
}

type edge struct {
	Permission        githubv4.String
	PermissionSources struct {
		Source githubv4.String
	}
	Node struct {
		DatabaseID githubv4.Int
		Name       githubv4.String
		Login      githubv4.String
		Email      githubv4.String
	}
}

// Collaborator struct
type Collaborator struct {
	ID    int
	Name  string
	Login string
	Email string
	Role  string
}

// New Users
func New(owner string, name string) *Collaborators {
	return &Collaborators{
		owner: owner,
		name:  name,
	}
}

// SetGitHubV4Client aaa
func (c *Collaborators) SetGitHubV4Client(client upstream.GraphQL) {
	c.client = client
}

// Load load members
func (c *Collaborators) Load() {
	if len(c.collaborators) > 0 {
		return
	}
	var query struct {
		RateLimit  rateLimit
		Repository struct {
			ID            githubv4.String
			DatabaseID    githubv4.Int
			Name          githubv4.String
			Collaborators struct {
				TotalCount githubv4.Int
				PageInfo   pageInfo
				Edges      []edge
			} `graphql:"collaborators(affiliation: $affiliation, first: $first, after: $after)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner":       githubv4.String(c.owner),
		"name":        githubv4.String(c.name),
		"affiliation": githubv4.CollaboratorAffiliationAll,
		// "affiliation": githubv4.CollaboratorAffiliationDirect,
		// "affiliation": githubv4.CollaboratorAffiliationOutside,
		"first": githubv4.Int(100),
		"after": (*githubv4.String)(nil),
	}
	i := 0
	for {
		pp.Printf("%+v\n", variables)
		err := c.client.Query(&query, variables)
		if err != nil {
			pp.Printf("%+v\n", err)
			os.Exit(1)
		}
		pp.Printf("%+\n", query)
		if !query.Repository.Collaborators.PageInfo.HasPreviousPage {
			size := query.Repository.Collaborators.TotalCount
			c.collaborators = make([]Collaborator, size, size)
		}
		for _, e := range query.Repository.Collaborators.Edges {
			c.collaborators[i] = Collaborator{
				ID:    int(e.Node.DatabaseID),
				Name:  string(e.Node.Name),
				Login: string(e.Node.Login),
				Email: string(e.Node.Email),
			}
			i++
		}
		pp.Printf("%+v", c.collaborators)
		if !query.Repository.Collaborators.PageInfo.HasNextPage {
			break
		}
		variables["after"] = githubv4.String(query.Repository.Collaborators.PageInfo.EndCursor)
	}
}
