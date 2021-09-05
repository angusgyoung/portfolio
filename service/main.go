package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    router.GET("/ping", ping)
	
    router.Run("localhost:8080")
}

type Response struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

func NewResponse(message string) *Response {
	return &Response{
		Message: message,
		Timestamp: time.Now().Unix(),
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, NewResponse("Available"))
}