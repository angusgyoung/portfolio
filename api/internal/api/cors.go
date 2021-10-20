package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

var CORS_ALLOW_ORIGIN = os.Getenv("CORS_ALLOW_ORIGIN")

func CORS() gin.HandlerFunc {

    var corsAllowedOrigin = "http://localhost:8080"
    if CORS_ALLOW_ORIGIN != "" {
        corsAllowedOrigin = CORS_ALLOW_ORIGIN
    }

    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", corsAllowedOrigin)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}