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

	github.Attach(v1)
	
    router.Run("localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResponse("Available"))
}