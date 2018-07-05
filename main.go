package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Initialise logging
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)

	log.Info("Server Started")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	fmt.Fprintf(w, `{"Message":"Hey"}`)
}

func logRequest(req *http.Request) {
	requestLogger := log.WithFields(log.Fields{
		"method":   req.Method,
		"req_addr": req.RemoteAddr,
		"url":      req.URL.Path,
	})

	requestLogger.Info("New Request")
}
