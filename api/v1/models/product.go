package models

import (
	"encoding/json"
	"fmt"
	"hash/fnv"

	"github.com/pkg/errors"
)

// Products is a collection of objects
type Products []Product

// Product is an item within a store.
type Product struct {
	ID    uint32  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Desc  string  `json:"desc,omitempty"`
	Price float64 `json:"price,omitempty"`
}

// NewProduct is a factory for creating new products unsurprisingly
func NewProduct(name, desc string, price float64) *Product {
	id := generateID(name, desc, price)
	return &Product{
		ID:    id,
		Name:  name,
		Desc:  desc,
		Price: price,
	}
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
// will need to update this in future if relational database is used
func generateID(name, desc string, price float64) uint32 {
	concat := fmt.Sprintf("%s%s%.2f", name, desc, price)

	hash := fnv.New32a()
	hash.Write([]byte(concat))

	return hash.Sum32()
}
