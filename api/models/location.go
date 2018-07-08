package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Tax describes a specific type of tax
type Tax struct {
	Name   string
	Amount float64
	Total  float64
}

// Currency describes a specific currency type
// refer to https://godoc.org/golang.org/x/text/currency#Unit for definition of Units
// The name should adhere to https://en.wikipedia.org/wiki/ISO_4217
type Currency struct {
	Name        string
	CountryName string
	Symbol      string
}

// Location is a specific geographic area with specific tax laws
// Generally will refer to the location that a transaction is carried out in
type Location struct {
	Name     string
	Currency Currency
	Taxes    []Tax
}

// NewLocation is a factory for creating new locations unsurprisingly
func NewLocation(name, currency string, taxes []Tax) (*Location, error) {
	c, err := getCurrency(currency)
	if err != nil {
		return &Location{}, errors.Wrap(err, "Issue Validating Currency")
	}
	return &Location{
		Name:     name,
		Currency: c,
		Taxes:    taxes,
	}, nil
}

// GetLocation information for a given location name
func GetLocation(name string) (*Location, error) {
	for _, loc := range getKnownLocations() {
		if name == loc.Name {
			return &loc, nil
		}
	}
	return &Location{}, errors.Errorf("Location Not Found: %s", name)
}

// getCurrency ensures that the currency name provided exists
func getCurrency(name string) (Currency, error) {
	for _, curr := range getKnownCurrencies() {
		if name == curr.Name {
			return curr, nil
		}
	}
	return Currency{}, errors.Errorf("Currency Not Found: %s", name)
}

// GetTaxes is a getter method to get the associated taxes in a location
func (l *Location) GetTaxes() ([]Tax, error) {
	taxes := l.Taxes
	if len(taxes) < 1 {
		return taxes, errors.New("No Taxes associated with Location")
	}
	return taxes, nil
}

// JSON returns a JSON representation of the location
func (l *Location) JSON() ([]byte, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Region Object")
	}

	return b, nil
}
