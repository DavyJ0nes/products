package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felixge/httpsnoop"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/davyj0nes/products/api/models"
	"github.com/gorilla/mux"
)

// Error describes the JSON response for a non 200
type Error struct {
	Message string `json:"message,omitempty"`
}

var (
	httpRequestsResponseTime prometheus.Summary
	requestDuration          *prometheus.HistogramVec
)

func init() {
	// initialise request duration metric
	requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Time (in secods) spent serving HTTP requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route", "status_code"})

	prometheus.MustRegister(requestDuration)
}

// Router is the mux Router for the Service
func Router(buildTime, commit, release string) http.Handler {
	// Setting up in memeory data
	models.Seed()

	r := mux.NewRouter()

	// basic endpoints
	r.HandleFunc("/version", version(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")

	// product endpoints
	// r.HandleFunc("/api/v1/product", newProduct).Methods("POST")
	r.HandleFunc("/api/v1/product/all", allProducts).Methods("GET")
	r.HandleFunc("/api/v1/product/{sku}", getProduct).Methods("GET")

	// transaction endpoints
	r.HandleFunc("/api/v1/transaction", newTransaction).Methods("POST")
	r.HandleFunc("/api/v1/transaction/all", allTransactions).Methods("GET")
	r.HandleFunc("/api/v1/transaction/{id}", getTransaction).Methods("GET")

	// prometheus metrics endpoint
	r.Handle("/metrics", promhttp.Handler())

	withMetrics := middleware(r)
	return withMetrics
}

// generateJSONResponse is a helper to allow for easier output of JSON data
func generateJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// middleware runs a set of functions on every request.
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			m := httpsnoop.CaptureMetrics(next, w, r)

			requestDuration.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(m.Code)).Observe(m.Duration.Seconds())
		})
}
