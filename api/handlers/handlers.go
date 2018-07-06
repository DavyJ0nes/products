package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/davyj0nes/products/api/v1/models"
	"github.com/gorilla/mux"
)

// Error describes the JSON response for a non 200
type Error struct {
	Message string `json:"message,omitempty"`
}

// Router is the mux Router for the Service
func Router(buildTime, commit, release string) *mux.Router {
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

	return r
}

// generateJSONResponse is a helper to allow for easier output of JSON data
func generateJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
