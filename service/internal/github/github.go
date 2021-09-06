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

func GetProfileData() (interface{}, error) {
    request := createRequest(userDataQuery)
	return executeRequest(*request)
}

func GetProjectData() (interface{}, error) {
    request := createRequest(projectDataQuery)
	return executeRequest(*request)
}

func executeRequest(request graphql.Request) (interface{}, error) {
	var response interface{}
    
	err := client.Run(context.Background(), &request, &response)

	if err != nil {
        log.WithError(err).Error("Failed to perform github request")
		return nil, err
    }
    
	log.Trace(response)
	return response, nil
}

func createRequest(query string) *graphql.Request {
	request := graphql.NewRequest(query)

	log.Trace(request)

	request.Header.Add("Authorization", "Bearer " + apiToken)

	return request
}

