package main

import (
	"flag"

	"github.com/angusgyoung/portfolio-service/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	logLevel := flag.String("logLevel", "warn", "interal logging level")
    flag.Parse()

	level, err := log.ParseLevel(*logLevel)

	if err != nil {
		log.WithError(err).Fatal("Failed to parse log level from flag")
	}

	log.SetLevel(level)

	internal.Init()	
}
