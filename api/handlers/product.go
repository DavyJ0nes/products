package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davyj0nes/products/api/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NewProductInput is used to parse the JSON input to newProduct
type NewProductInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Colour      string `json:"colour,omitempty"`
	SKU         string `json:"sku,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Price       int    `json:"price,omitempty"`
}

// NewProductOutput is used as the JSON output for newProduct
type NewProductOutput struct {
	ID   uint32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	log.Info("Received Request: ", "allProduct")
	products := models.KnownProducts
	generateJSONResponse(w, http.StatusOK, products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Info("Received Request: ", "getProduct")
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

func newProduct(w http.ResponseWriter, r *http.Request) {
	log.Info("Received Request: ", "newProduct")
	var (
		newProdInput  NewProductInput
		newProdOutput NewProductOutput
	)

	if r.Body == nil {
		errMsg := Error{"received no data"}
		generateJSONResponse(w, http.StatusBadRequest, errMsg)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		checkError(w, err)
		return
	}

	if len(b) == 0 {
		errMsg := Error{"received no data"}
		generateJSONResponse(w, http.StatusBadRequest, errMsg)
		return
	}
	err = json.Unmarshal(b, &newProdInput)
	if err != nil {
		checkError(w, err)
		return
	}

	// Create new product instance
	prod := models.NewProduct(newProdInput.Name, newProdInput.Description, newProdInput.Colour, newProdInput.SKU, newProdInput.Currency, newProdInput.Price)

	// Store product in data store
	err = models.StoreProduct(prod)
	if err != nil {
		checkError(w, err)
		return
	}

	log.Infof("Created New Product: %v", prod.ID)

	// Generate return Output
	newProdOutput.ID = prod.ID
	newProdOutput.Name = prod.Name
	generateJSONResponse(w, http.StatusOK, newProdOutput)
}
