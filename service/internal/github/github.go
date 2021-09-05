package github

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	log "github.com/sirupsen/logrus"
)

var apiUrl = os.Getenv("GITHUB_API_URL")
var apiToken = os.Getenv("GITHUB_API_TOKEN")

var client graphql.Client = *graphql.NewClient(apiUrl)

func Attach(base *gin.RouterGroup) {
	github := base.Group("/github")

	github.GET("/user", getProfileData)
	github.GET("/projects", getProjectData)
}

func getProfileData(ctx *gin.Context) {
    request := createRequest(userDataQuery)

    var response interface{}
    
	if err := client.Run(context.Background(), request, &response); err != nil {
        log.Fatal(err)
    }
    
	log.Trace(response)
	ctx.JSON(http.StatusOK, response)
}

func getProjectData(ctx *gin.Context) {
    request := createRequest(projectDataQuery)

    var response interface{}
    
	if err := client.Run(context.Background(), request, &response); err != nil {
        panic(err)
    }
    
	log.Trace(response)
	ctx.JSON(http.StatusOK, response)
}

func createRequest(query string) *graphql.Request {
	request := graphql.NewRequest(query)

	log.Trace(request)

	request.Header.Add("Authorization", "Bearer " + apiToken)

	return request
}

