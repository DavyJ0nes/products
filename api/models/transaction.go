package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/satori/go.uuid"

	"github.com/pkg/errors"
)

// Transactions is a collection of objects
type Transactions []Transaction

// Transaction is a representation of a payment transaction
// This could be an invoice, bill etc
type Transaction struct {
	ID            string    `json:"id,omitempty"`
	Datetime      time.Time `json:"datetime,omitempty"`
	Location      *Location `json:"location,omitempty"`
	ConversionAPI string
	Products      []Product `json:"products,omitempty"`
	Subtotal      int       `json:"subtotal,omitempty"`
	TaxTotal      int       `json:"tax_total,omitempty"`
	Total         int       `json:"total,omitempty"`
}

// NewTransaction starts a new Transaction
// All Transactions must specify a region as this is used to get the current price of the products
func NewTransaction(location, conversionAPIURL string) (*Transaction, error) {
	id, err := generateTransactionID()
	if err != nil {
		return &Transaction{}, errors.Wrap(err, "Problem generating UUID for Transaction")
	}

	l, err := GetLocation(location)
	if err != nil {
		return &Transaction{}, errors.Wrap(err, "Problem Getting Location Information")
	}

	return &Transaction{
		ID:            id,
		Datetime:      time.Now(),
		Location:      l,
		ConversionAPI: conversionAPIURL,
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
	var runningTotal int

	for _, product := range t.Products {
		conversionRate, err := t.getLocalRate(product.BaseCurrency, t.Location.Currency.Name)
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

	for idx, tax := range t.Location.Taxes {
		runningTotal += float64(t.Subtotal) * tax.Amount
		// Updating the pointer to the Tax
		// tax is a copy not pointer, which is why need to do this
		t.Location.Taxes[idx].Total = int(float64(t.Subtotal) * tax.Amount)
	}

	t.TaxTotal = int(runningTotal)

}

// GetTaxBreakdown gives breakdown of each of the taxes in a transaction
// requires CalcTaxTotal() to have been run before to ensure Tax Totals != 0.0
func (t *Transaction) GetTaxBreakdown() []Tax {
	for i, tax := range t.Location.Taxes {
		// update tax total
		// TODO (davy): Not 100% sure if this is a good idea
		if tax.Total == 0 {
			t.Location.Taxes[i].Total = int(float64(t.Subtotal) * tax.Amount)
		}
	}

	return t.Location.Taxes
}

// CalcTransactionTotal creates the final total for the transaction
// This will be then need to be paid
func (t *Transaction) CalcTransactionTotal() {
	t.Total = t.Subtotal + t.TaxTotal
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
// TODO (davy): Will require refactoring when database is implemented
func generateTransactionID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// rateResponse is used when parsing the output from the currency convertor
type rateResponse struct {
	Type struct {
		Val float64 `json:"val,omitempty"`
	} `json:"type,omitempty"`
}

// calcLocalPrice queries a free currency convertor to get an up to date rate for
// the base currency and the locations currency
// This is a simple method of ensuring that the price given is correct.
// This method is pretty horrible and in need of refactoring
// For production use this should be its own package to handle edge cases
func calcLocalPrice(basePrice int, rate float64) int {
	localPrice := float64(basePrice) * rate

	return int(localPrice)

}

type convertorErrorRepsonse struct {
	Status int
	Error  string
}

// getLocalRate queries a free currency convertor to get an up to date rate for
// the base currency and the locations currency
// This API has a rate limit of 100 requests per hour
// more info: https://free.currencyconverterapi.com/
func (t *Transaction) getLocalRate(baseCurrency, locationCurrency string) (float64, error) {
	queryKey := fmt.Sprintf("%s_%s", baseCurrency, locationCurrency)

	// Query is FROM currency TO Currency
	URL := fmt.Sprintf("%s?q=%s&compact=y", t.ConversionAPI, queryKey)
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

	// due to not knowing what the key of the currency types am having to take care of it
	// the response structure looks like: {"GBP_USD":{"val":1.321406}}
	pattern := regexp.MustCompile("[A-Z]{3}_[A-Z]{3}")
	parsedBody := pattern.ReplaceAllString(string(body), "type")

	var data rateResponse
	err = json.Unmarshal([]byte(parsedBody), &data)
	if err != nil {
		return 0.0, err
	}

	return data.Type.Val, nil
}
