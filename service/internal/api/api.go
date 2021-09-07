package api

import (
	"net/http"

	"github.com/angusgyoung/portfolio-service/internal/cache"
	"github.com/angusgyoung/portfolio-service/internal/github"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(ApplicationErrorHandler())
	
	v1 := router.Group("/api/v1")
	
	v1.GET("/ping", ping)
	
	ghub := v1.Group("/github")
	
	ghub.GET("/user", getProfileData)
	ghub.GET("/projects", getProjectData)	
	
	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, NewResponse("Available"))
}

func getProfileData(c *gin.Context) {
	cache.PerformCacheableServiceCall(c, github.GetProfileData, "profileData")
}
	
func getProjectData(c *gin.Context) {
	cache.PerformCacheableServiceCall(c, github.GetProjectData, "projectData")
}

