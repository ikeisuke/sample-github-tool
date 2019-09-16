package main

import (
	"os"

	"github.com/ikeisuke/sample-github-tool/github"
	"github.com/k0kubun/pp"
)

func main() {
	g := github.New()
	o := g.Organization(os.Getenv("GITHUB_ORGANIZATION_NAME"))
	pp.Printf("%+v\n", o)
	m := o.Members()
	// m.Load()
	pp.Printf("%+v\n", m)
	r := g.Repository(os.Getenv("GITHUB_ORGANIZATION_NAME"), os.Getenv("GITHUB_REPOSITORY_NAME"))
	pp.Printf("%+v\n", r)
	c := r.Collaborators()
	// c.Load()
	pp.Printf("%+v\n", c)
	// r.CreateIssue("test title", "test body \n new line")
}
