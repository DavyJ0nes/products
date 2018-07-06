package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/davyj0nes/products/api/handlers"
	"github.com/davyj0nes/products/api/version"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Initialise logging
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// CONFIGURING COMMAND LINE ARGS
	port := flag.String("port", "8080", "port number to use with Web Service")
	flag.Parse()

	log.Printf("Starting Service... | {commit: %s, build_time: %s, release: %s}", version.Commit, version.BuildTime, version.Release)
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
