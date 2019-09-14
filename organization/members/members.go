package members

import (
	"context"
	"os"

	"github.com/k0kubun/pp"
	"github.com/shurcooL/githubv4"
)

// Members struct
type Members struct {
	client       *githubv4.Client
	organization string
	members      []Member
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
	Role githubv4.String
	Node struct {
		DatabaseId githubv4.Int
		Name       githubv4.String
		Login      githubv4.String
		Email      githubv4.String
	}
}

// Member struct
type Member struct {
	ID    int
	Name  string
	Login string
	Email string
	Role  string
}

// New Users
func New(organization string) *Members {
	return &Members{
		organization: organization,
	}
}

// SetGitHubV4Client aaa
func (m *Members) SetGitHubV4Client(client *githubv4.Client) {
	m.client = client
}

// Load load members
func (m *Members) Load() {
	if len(m.members) > 0 {
		return
	}
	var query struct {
		RateLimit    rateLimit
		Organization struct {
			Login           githubv4.String
			Name            githubv4.String
			MembersWithRole struct {
				TotalCount githubv4.Int
				PageInfo   pageInfo
				Edges      []edge
			} `graphql:"membersWithRole(first: $first, after: $after)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login": githubv4.String(m.organization),
		"first": githubv4.Int(100),
		"after": (*githubv4.String)(nil),
	}
	i := 0
	for {
		err := m.client.Query(context.Background(), &query, variables)
		if err != nil {
			pp.Printf("%+v\n", err)
			os.Exit(1)
		}
		//pp.Printf("%+v\n", query)
		if !query.Organization.MembersWithRole.PageInfo.HasPreviousPage {
			size := query.Organization.MembersWithRole.TotalCount
			m.members = make([]Member, size, size)
		}
		for _, e := range query.Organization.MembersWithRole.Edges {
			m.members[i] = Member{
				ID:    int(e.Node.DatabaseId),
				Name:  string(e.Node.Name),
				Login: string(e.Node.Login),
				Email: string(e.Node.Email),
				Role:  string(e.Role),
			}
			i++
		}
		if !query.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}
		variables["after"] = githubv4.String(query.Organization.MembersWithRole.PageInfo.EndCursor)
	}
}
