package handlers

import (
	"github.com/gorilla/mux"
)

type Error struct {
	Message string `json:"message,omitempty"`
}

// Router is the mux Router for the Service
func Router(buildTime, commit, release string) *mux.Router {
	r := mux.NewRouter()

	// basic endpoints
	r.HandleFunc("/version", version(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")

	// product endpoints
	r.HandleFunc("/product", newProduct).Methods("POST")
	r.HandleFunc("/product/all", allProducts).Methods("GET")
	r.HandleFunc("/product/{id}", getProduct).Methods("GET")

	// transaction endpoints
	r.HandleFunc("/transaction", newTransaction).Methods("POST")
	r.HandleFunc("/transaction/all", allTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", getTransaction).Methods("GET")

	return r
}
