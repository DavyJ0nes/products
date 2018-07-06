package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/satori/go.uuid"

	"github.com/pkg/errors"
)

// Transactions is a collection of objects
type Transactions []Transaction

// Transaction is a representation of a payment transaction
// This could be an invoice, bill etc
type Transaction struct {
	ID       string    `json:"id,omitempty"`
	Datetime time.Time `json:"datetime,omitempty"`
	Location *Location `json:"location,omitempty"`
	Products []Product `json:"products,omitempty"`
	Subtotal float64   `json:"subtotal,omitempty"`
	TaxTotal float64   `json:"tax_total,omitempty"`
	Total    float64   `json:"total,omitempty"`
}

// NewTransaction starts a new Transaction
// All Transactions must specify a region as this is used to get the current price of the products
func NewTransaction(location string) (*Transaction, error) {
	id, err := generateTransactionID()
	if err != nil {
		return &Transaction{}, errors.Wrap(err, "Problem generating UUID for Transaction")
	}

	l, err := GetLocation(location)
	if err != nil {
		return &Transaction{}, errors.Wrap(err, "Problem Getting Location Information")
	}

	return &Transaction{
		ID:       id,
		Datetime: time.Now(),
		Location: l,
	}, nil
}

// StoreTransaction adds a completed transaction to the data store
// This will need to be refactored once data store has been added
// TODO (davy): Refactor for datastore
func StoreTransaction(tran *Transaction) error {
	KnownTransactions = append(KnownTransactions, *tran)

	return nil
}

// AddProducts is used to add multiple products to a transaction
func (t *Transaction) AddProducts(products []Product) {
	for _, p := range products {
		t.addProduct(p)
	}
}

// addProduct is used to add a new Product to the transactions Line Items
func (t *Transaction) addProduct(product Product) {
	t.Products = append(t.Products, product)
}

// CalcSubtotal totals the prices of each of the Products in the transaction
// in the locations Currency
func (t *Transaction) CalcSubtotal() error {
	var runningTotal float64

	for _, product := range t.Products {
		conversionRate, err := getLocalRate(product.BaseCurrency, t.Location.Currency.Name)
		if err != nil {
			return errors.Wrap(err, "Problem getting conversion rate")
		}

		localPrice := calcLocalPrice(product.BasePrice, conversionRate)
		// Update local price. Used for transaction output
		product.LocalPrice = localPrice
		runningTotal += localPrice

	}
	t.Subtotal = runningTotal

	return nil
}

// CalcTaxTotal iterates through the locations taxes and applies them to the subtotal
func (t *Transaction) CalcTaxTotal() {
	var runningTotal float64

	for _, tax := range t.Location.Taxes {
		runningTotal += (t.Subtotal * tax.Amount)
	}

	t.TaxTotal = formatAmount(runningTotal)

}

// CalcTransactionTotal creates the final total for the transaction
// This will be then need to be paid
func (t *Transaction) CalcTransactionTotal() {
	total := t.Subtotal + t.TaxTotal
	t.Total = formatAmount(total)
}

// JSON returns a JSON representation of the transaction
func (t *Transaction) JSON() ([]byte, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Transaction Object")
	}

	return b, nil
}

// generateTransactionID is used to create a UUID for a transaction
// will require refactoring when database is implemented
func generateTransactionID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// rateVal is used when parsing the output from the currency convertor
type rateVal struct {
	Val float64 `json:"val,omitempty"`
}

// calcLocalPrice queries a free currency convertor to get an up to date rate for
// the base currency and the locations currency
// This is a simple method of ensuring that the price given is correct.
// This method is pretty horrible and in need of refactoring
// For production use this should be its own package to handle edge cases
func calcLocalPrice(basePrice, rate float64) float64 {
	// Calculate the local price by multiplying a rounded conversion rate
	precision := math.Pow(10, float64(2))
	localPrice := basePrice * math.Round(rate*precision) / precision

	// This is horrible but couldn't find a better way of getting output to 2 decimals.
	// TODO (davy): Find better way of handling this (seperate package)
	formattedPrice, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", localPrice), 64)

	return formattedPrice

}

type convertorErrorRepsonse struct {
	Status int
	Error  string
}

// getLocalRate queries a free currency convertor to get an up to date rate for
// the base currency and the locations currency
// This API has a rate limit of 100 requests per hour
// more info: https://free.currencyconverterapi.com/
// TODO (davy): mock conversion service
func getLocalRate(baseCurrency, locationCurrency string) (float64, error) {
	baseURL := "http://free.currencyconverterapi.com/api/v5/convert"
	queryKey := fmt.Sprintf("%s_%s", baseCurrency, locationCurrency)

	// Query is FROM currency TO Currency
	URL := fmt.Sprintf("%s?q=%s&compact=y", baseURL, queryKey)
	response, err := http.Get(URL)
	if err != nil {
		return 0.0, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}

	if response.StatusCode != http.StatusOK {
		var errMsg convertorErrorRepsonse
		json.Unmarshal(body, &errMsg)
		return 0.0, errors.Errorf("Non 200 Response: %d\n%s", errMsg.Status, errMsg.Error)
	}

	// Parse body as Map
	var m map[string]rateVal
	err = json.Unmarshal(body, &m)
	if err != nil {
		return 0.0, err
	}

	// Calculate the local price by multiplying a rounded conversion rate
	rateVal := m[queryKey].Val

	return rateVal, nil
}

func formatAmount(rawAmount float64) float64 {
	// Rounding to 2 decimal places. Is a bit of a hack for now
	// TODO (davy): Find better way of handling this (seperate package)
	precision := math.Pow(10, float64(2))
	preFormattedAmount := math.Round(rawAmount*precision) / precision
	formattedAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", preFormattedAmount), 64)

	return formattedAmount
}
