package internal

import (
	"net/http"

	"github.com/angusgyoung/portfolio-service/internal/github"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

    v1.GET("/ping", ping)

	ghub := v1.Group("/github")

	ghub.GET("/user", getProfileData)
	ghub.GET("/projects", getProjectData)	
	
    router.Run("localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResponse("Available"))
}

func getProfileData(ctx *gin.Context) {
	data := github.GetProfileData()

	ctx.JSON(http.StatusOK, data)
}

func getProjectData(ctx *gin.Context) {
	data := github.GetProfileData()

	ctx.JSON(http.StatusOK, data)
}
