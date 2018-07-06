package handlers

import (
	"fmt"
	"net/http"

	"github.com/davyj0nes/products/api/v1/models"
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
	if err != nil {
		errMessage := Error{
			Message: fmt.Sprintf("Problem finding product: %v", err),
		}
		generateJSONResponse(w, http.StatusNotFound, errMessage)
		return
	}

	generateJSONResponse(w, http.StatusOK, product)
}
