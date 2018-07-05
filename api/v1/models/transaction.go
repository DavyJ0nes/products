package models

import (
	"encoding/json"
	"time"

	"github.com/satori/go.uuid"

	"github.com/davyj0nes/products-api/product"
	"github.com/pkg/errors"
)

// Transaction is a representation of a payment transaction
// This could be an invoice, bill etc
type Transaction struct {
	ID         string            `json:"id,omitempty"`
	Datetime   time.Time         `json:"datetime,omitempty"`
	RegionName string            `json:"region_name,omitempty"`
	Products   []product.Product `json:"products,omitempty"`
	Subtotal   float64           `json:"subtotal,omitempty"`
	TaxTotal   float64           `json:"tax_total,omitempty"`
	Total      float64           `json:"total,omitempty"`
}

// NewTransaction instanciates a Transaction
func NewTransaction() (*Transaction, error) {
	id, err := generateID()
	if err != nil {
		return &Transaction{}, errors.Wrap(err, "Problem generating UUID for Transaction")
	}

	return &Transaction{
		ID:       id,
		Datetime: time.Now(),
	}, nil
}

func (t *Transaction) AddProduct(product product.Product) {
	t.Products = append(t.Products, product)
}

func (t *Transaction) CalcSubtotal() {
	var runningTotal float64

	for _, product := range t.Products {
		runningTotal += product.Price

	}
	t.Subtotal = runningTotal
}

func (t *Transaction) CalcTaxTotal() error {
	t.TaxTotal = 0.0
	return nil
}

func (t *Transaction) CalcTransactionTotal() error {
	t.Total = t.Subtotal + t.TaxTotal
	return nil
}

// JSON returns a JSON representation of the transaction
func (t *Transaction) JSON() ([]byte, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Transaction Object")
	}

	return b, nil
}

// generateID is used to create a UUID for a transaction
// will require refactoring when database is implemented
func generateID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
