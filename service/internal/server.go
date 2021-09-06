package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/angusgyoung/portfolio-service/internal/github"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var cache *redis.Client

var CACHE_TTL = os.Getenv("CACHE_TTL")
var CACHE_URL = os.Getenv("CACHE_URL")

func Init() {
	cache = redis.NewClient(&redis.Options{
		Addr:     CACHE_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	
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
	performCacheableServiceCall(c, github.GetProfileData, "profileData")
}
	
func getProjectData(c *gin.Context) {
	performCacheableServiceCall(c, github.GetProjectData, "projectData")
}

func performCacheableServiceCall(c *gin.Context, retrieveCallback func() (interface{}, error), cacheKey string) {
	var data interface{}

	log.WithField("cacheKey", cacheKey).Trace("Checking for cached response")
	cached, err := cache.Get(ctx, cacheKey).Result()

	if err == redis.Nil {
		log.Debug("Response is not cached, retrieving")
		data, err = retrieveCallback()

		if err != nil {
			c.Error(err)
			return
		}

		jsonString, err := json.Marshal(data)

		if err != nil {
			log.WithError(err).Error("Failed to marshall response to string for caching")
		}

		// default TTL to an hour
		cacheTtl := 3600
		if CACHE_TTL != "" {
			cacheTtl, _ = strconv.Atoi(CACHE_TTL)
		}

		log.WithField("cacheTtl", cacheTtl).Trace("Caching response")
		err = cache.Set(ctx, cacheKey, jsonString,  time.Duration(cacheTtl) * time.Second).Err()

		if err != nil {
			log.WithError(err).Error("Failed to cache response")
		}
		
		log.Trace("Executed retrieve callback")
	} else if err != nil {
		c.Error(err)
		return		
	} else if err != redis.Nil {
		log.Debug("Response exists in cache")
		json.Unmarshal([]byte(cached), &data)
	}

	c.JSON(http.StatusOK, data)
}
		