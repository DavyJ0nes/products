package handlers

import (
	"encoding/json"
	"net/http"
)

func newProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// vars := mux.Vars(r)
	// id := vars["id"]
	errMessage := Error{"User not found"}
	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusNotFound)

	encoder.Encode(errMessage)

}
