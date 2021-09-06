package internal

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type applicationError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (e *applicationError) Error() string {
	return e.Message
}

func ApplicationErrorHandler() gin.HandlerFunc {
    return applicationErrorHandler(gin.ErrorTypeAny)
}

func applicationErrorHandler(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
        c.Next()
        detectedErrors := c.Errors.ByType(errType)
        
		if len(detectedErrors) > 0 {
            err := detectedErrors[0].Err

			log.WithError(err).Trace("Handling application error")

			var parsedError *applicationError
            
			switch err := err.(type) {
            case *applicationError:
                parsedError = err
            default:
                parsedError = &applicationError{ 
                  Code: http.StatusInternalServerError,
                  Message: "Internal Server Error",
                }
            }
            
			c.AbortWithStatusJSON(parsedError.Code, parsedError)
            return
        }
	}
}