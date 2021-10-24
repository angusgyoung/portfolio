package cache

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var cache *redis.Client

var CACHE_TTL = os.Getenv("CACHE_TTL")
var CACHE_URL = os.Getenv("CACHE_URL")
var CACHE_TIMEOUT = os.Getenv("CACHE_TIMEOUT")

func Init() {
	cacheUrl := "localhost:6379"
	if CACHE_URL != "" {
		cacheUrl = CACHE_URL
	}

	var cacheTimeout int
	if CACHE_TIMEOUT == "" {
		cacheTimeout = 3
	} else {
		cacheTimeout, _ = strconv.Atoi(CACHE_TIMEOUT)
	}

	options := &redis.Options{
		Addr:     cacheUrl,
		Password: "",
		DB:       0,
		DialTimeout: toSecondDuration(cacheTimeout),
		ReadTimeout: toSecondDuration(cacheTimeout),
		MaxRetries: -1,
		OnConnect: onConnect,
	}

	log.WithField("options", options).Trace("Creating Redis client")

	cache = redis.NewClient(options)
}

func PerformCacheableServiceCall(c *gin.Context, retrieveCallback func() (interface{}, error), cacheKey string) {
	var data interface{}

	log.WithField("cacheKey", cacheKey).Trace("Checking for cached response")
	cached, err := cache.Get(ctx, cacheKey).Result()

	if err != nil {

		if err != redis.Nil {
			// We have encountered an error attempting to retrieve something
			// from the cache. We should warn of the error but fallback to calling
			// the retrieveCallback
			log.WithError(err).Warn("Failed to retrieve value from cache")
		} else {
			// There is nothing in the cache for the key
			log.Debug("Response is not cached")
		}

		data, err = retrieveCallback()

		if err != nil {
			c.Error(err)
			return
		}

		jsonString, err := json.Marshal(data)

		if err != nil {
			log.WithError(err).Warn("Failed to marshall response to string for caching")
		} else {
			// default TTL to an hour
			cacheTtl := 3600
			if CACHE_TTL != "" {
				cacheTtl, _ = strconv.Atoi(CACHE_TTL)
			}

			log.WithField("cacheTtl", cacheTtl).Trace("Caching response")
			err = cache.Set(ctx, cacheKey, jsonString,  time.Duration(cacheTtl) * time.Second).Err()

			if err != nil {
				// Warn that the response couldn't be cached
				log.WithError(err).Warn("Failed to cache response")
			}
		}
		
		log.Trace("Executed retrieve callback")
	} else {
		log.Debug("Returning cached response")
		json.Unmarshal([]byte(cached), &data)
	}

	c.JSON(http.StatusOK, data)
}

func onConnect(ctx context.Context, cn *redis.Conn) error {
	log.WithFields(log.Fields{
		"ClientID": cn.ClientID(ctx),
		"Version": redis.Version(),
	}).Debug("Connected to redis")
	return nil
}

func toSecondDuration(seconds int) time.Duration {
	return time.Duration(seconds * int(time.Second))
}