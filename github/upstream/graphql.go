package upstream

import "github.com/shurcooL/githubv4"

type GraphQL interface {
	Query(q interface{}, variables map[string]interface{}) error
	Mutate(m interface{}, input githubv4.Input, variables map[string]interface{}) error
}
