package github

import (
	"context"
	"os"

	"github.com/hasura/go-graphql-client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var apiUrl = os.Getenv("GITHUB_API_URL")
var apiToken = os.Getenv("GITHUB_API_TOKEN")

var src = oauth2.StaticTokenSource(
	&oauth2.Token{
		AccessToken: apiToken,
		TokenType:   "Bearer",
	},
)

var httpClient = oauth2.NewClient(context.Background(), src)
var client graphql.Client = *graphql.NewClient(apiUrl, httpClient)


func GetProfileData() (interface{}, error) {
    return performRequest(&userDataQuery)
}

func GetProjectData() (interface{}, error) {
    return performRequest(&projectDataQuery)
}

func performRequest(q interface{}) (interface{}, error) {
	err := client.Query(context.Background(), q, nil)
	log.WithFields(log.Fields{
		"err": err,
		"response": q,
	}).Trace("Executed query")
	return q, err
}


