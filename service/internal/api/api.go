package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/angusgyoung/portfolio-service/internal/cache"
	"github.com/angusgyoung/portfolio-service/internal/github"
	"github.com/gin-gonic/gin"
)

var API_PORT = os.Getenv("API_PORT")
var buildVersion string

func Init(version string) {
	buildVersion = version

	router := gin.Default()
	router.Use(ApplicationErrorHandler())
	
	v1 := router.Group("/api/v1")
	
	v1.GET("/ping", ping)
	
	ghub := v1.Group("/github")
	
	ghub.GET("/user", getProfileData)
	ghub.GET("/projects", getProjectData)

	apiPort := 8080
	if API_PORT != "" {
		var err error
		apiPort, err = strconv.Atoi(API_PORT)
		if err != nil {
			log.WithField("API_PORT", API_PORT).Fatal("Could not interpret a port number from the provided argument")
		}
	}
	
	router.Run(fmt.Sprint("localhost:", apiPort))
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, NewResponse("Available", buildVersion))
}

func getProfileData(c *gin.Context) {
	cache.PerformCacheableServiceCall(c, github.GetProfileData, "profileData")
}
	
func getProjectData(c *gin.Context) {
	cache.PerformCacheableServiceCall(c, github.GetProjectData, "projectData")
}

