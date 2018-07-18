package models

import (
	"encoding/json"
	"fmt"
	"hash/fnv"

	"github.com/pkg/errors"
)

// Products is a representation of a collection of Products
// Having to do this to ensure good JSON encoding
type Products struct {
	Products []Product `json:"products,omitempty"`
}

// Product is an item within a store.
type Product struct {
	ID           uint32 `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Desc         string `json:"desc,omitempty"`
	Colour       string `json:"colour,omitempty"`
	SKU          string `json:"sku,omitempty"`
	BasePrice    int    `json:"price,omitempty"`
	BaseCurrency string `json:"base_currency,omitempty"`
	LocalPrice   int    `json:"local_price,omitempty"`
}

// NewProduct is a factory for creating new products unsurprisingly
func NewProduct(name, desc, colour, sku, currency string, price int) *Product {
	id := generateProductID(name, desc, colour, sku)
	return &Product{
		ID:           id,
		Name:         name,
		Desc:         desc,
		Colour:       colour,
		SKU:          sku,
		BasePrice:    price,
		BaseCurrency: currency,
		LocalPrice:   price,
	}
}

// StoreProduct adds a newly created product to the in memory datastore
// This will need to be refactored once persistant data store has been added
// TODO (davy): Refactor for datastore
func StoreProduct(prod *Product) error {
	KnownProducts.Products = append(KnownProducts.Products, *prod)

	return nil
}

// GetProduct looks up an object by its SKU
func GetProduct(sku string) (*Product, error) {
	for _, product := range KnownProducts.Products {
		if product.SKU == sku {
			return &product, nil
		}
	}
	return &Product{}, errors.Errorf("No Product Matches SKU: %s", sku)
}

// GetProducts returns all known products
func GetProducts() []Product {
	// Currently just getting all in memory seeded Products
	// TODO (davy): Will need up update this once database is set up
	// KnownProducts is in seed.go
	return KnownProducts.Products
}

// JSON returns a JSON representation of all the products
func (p *Products) JSON() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Products Object")
	}

	return b, nil
}

// JSON returns a JSON representation of the product
func (p *Product) JSON() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Product Object")
	}

	return b, nil
}

// generateID creates a hash of the product information provided
// TODO (davy): Will need to update this in future if relational database is used
func generateProductID(name, desc, colour, sku string) uint32 {
	concat := fmt.Sprintf("%s%s%s%s", name, desc, colour, sku)

	hash := fnv.New32a()
	hash.Write([]byte(concat))

	return hash.Sum32()
}
