package handlers

import (
	"net/http"

	"github.com/davyj0nes/products/api/models"
	"github.com/gorilla/mux"
)

// func newProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// }

func allProducts(w http.ResponseWriter, r *http.Request) {
	products := models.KnownProducts
	generateJSONResponse(w, http.StatusOK, products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku := vars["sku"]
	product, err := models.GetProduct(sku)
	// having to do this due to execution continuing after error
	if err != nil {
		checkError(w, err)
		return
	}

	generateJSONResponse(w, http.StatusOK, product)
}
