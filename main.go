package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Create router with attached handlers
	log.Printf("Starting Service... | {commit: %s, build_time: %s, release: %s}", version.Commit, version.BuildTime, version.Release)
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	// srv object sets up prod grade http sever
	// for more info see: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts
	//   && https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
	srv := &http.Server{
		Addr:         ":" + *port,
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// create channel to handle sigterm signal for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Info("Service has started...")

	killSig := <-interrupt
	switch killSig {
	case os.Interrupt:
		log.Info("Received CTRL+C Interrupt Signal")

	case syscall.SIGTERM:
		log.Info("Received SIGTERM, Terminating...")
	}

	log.Info("Service shutting down ...")
	srv.Shutdown(context.Background())
	log.Info("Service has shut down")

}
