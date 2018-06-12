package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	githubpkg "github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"google.golang.org/appengine"
)

//  =================== REPOSITORY ===================

type repository struct {
	name, description, language, url string
	stars                            int
	archived                         bool
}

func (r repository) GetName() string {
	return r.name
}

func (r repository) GetDescription() string {
	return r.description
}

func (r repository) GetLanguage() string {
	return r.language
}

func (r repository) GetURL() string {
	return r.url
}

func (r repository) GetStars() int {
	return r.stars
}

func (r repository) IsArchived() bool {
	return r.archived
}

//  =================== GITHUB CLIENT ===================

type github struct {
	client
}

func (g *github) tokens() map[string]string {
	// Check if the default .env file must be used and if it exists
	if flag.Lookup("test.v") != nil {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal("Error loading testing .env file")
		}
	}

	return map[string]string{
		"GITHUB_OAUTH_TOKEN": os.Getenv("GITHUB_OAUTH_TOKEN"),
	}
}

func (g *github) getInfo(req *http.Request) []repository {

	ctx := appengine.NewContext(req)
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.tokens()["GITHUB_OAUTH_TOKEN"]},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := githubpkg.NewClient(tc)

	// list all repositories for the authenticated user
	content, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	repoLength := len(content)
	repoList := make([]repository, repoLength)
	for i, repo := range content {
		repoList[i] = repository{
			name:        repo.GetName(),
			description: repo.GetDescription(),
			language:    repo.GetLanguage(),
			stars:       repo.GetStargazersCount(),
			archived:    repo.GetArchived(),
			url:         repo.GetHTMLURL(),
		}
	}

	return repoList
}
