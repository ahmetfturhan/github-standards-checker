package githubFuncs

import (
	"context"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func Authorize(ACCESS_TOKEN, BASE_URL *string) (context.Context, *github.Client) {
	//Authorize with Access Token
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *ACCESS_TOKEN},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	//Create the GitHub Enterprise API Client
	client, err := github.NewEnterpriseClient(*BASE_URL, *BASE_URL, tokenClient)
	if err != nil {
		panic("Error while creating the Enterprise client")
	}
	return context, client
}
