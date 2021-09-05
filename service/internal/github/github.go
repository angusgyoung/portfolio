package github

import (
	"context"
	"os"

	"github.com/machinebox/graphql"
	log "github.com/sirupsen/logrus"
)

var apiUrl = os.Getenv("GITHUB_API_URL")
var apiToken = os.Getenv("GITHUB_API_TOKEN")

var client graphql.Client = *graphql.NewClient(apiUrl)

func GetProfileData() interface{} {
    request := createRequest(userDataQuery)

    var response interface{}
    
	if err := client.Run(context.Background(), request, &response); err != nil {
        log.Fatal(err)
    }
    
	log.Trace(response)

	return response
}

func GetProjectData() interface{} {
    request := createRequest(projectDataQuery)

    var response interface{}
    
	if err := client.Run(context.Background(), request, &response); err != nil {
        panic(err)
    }
    
	log.Trace(response)
	return response
}

func createRequest(query string) *graphql.Request {
	request := graphql.NewRequest(query)

	log.Trace(request)

	request.Header.Add("Authorization", "Bearer " + apiToken)

	return request
}

