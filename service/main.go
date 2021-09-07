package main

import (
	"flag"

	"github.com/angusgyoung/portfolio-service/internal/api"
	"github.com/angusgyoung/portfolio-service/internal/cache"
	log "github.com/sirupsen/logrus"
)

func main() {
	logLevel := flag.String("l", "warn", "internal logging level")
    flag.Parse()

	level, err := log.ParseLevel(*logLevel)

	if err != nil {
		log.WithError(err).Fatal("Failed to parse log level from flag")
	}

	log.SetLevel(level)

	log.Trace("Initialising cache")
	cache.Init()
	log.Trace("Initialising API")
	api.Init()	
}
