package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Tax describes a specific type of tax
type Tax struct {
	Name   string
	Amount float64
}

// Currency describes a specific currency type
// refer to https://godoc.org/golang.org/x/text/currency#Unit for definition of Units
type Currency struct {
	Name        string
	CountryName string
	Symbol      string
}

// Region is a specific geographic area with specific tax laws
// Generally will refer to the location that a transaction is carried out in
type Region struct {
	Name         string
	CurrencyName string
	Taxes        []Tax
}

// NewRegion is a factory for creating new regions unsurprisingly
func NewRegion(name, currency string, taxes []Tax) (*Region, error) {
	err := checkCurrency(currency)
	if err != nil {
		return &Region{}, errors.Wrap(err, "Issue Validating Currency")
	}
	return &Region{
		Name:         name,
		CurrencyName: currency,
		Taxes:        taxes,
	}, nil
}

// checkCurrency ensures that the currency name provided exists
func checkCurrency(name string) error {
	for _, curr := range getKnownCurrencies() {
		if name == curr.Name {
			return nil
		}
	}
	return errors.Errorf("Currency Not Found: %s", name)
}

// getKnownCurrencies looks for already defined currencies
// currently currencies are stored statically, this will need to be refactored once database has been added
func getKnownCurrencies() []Currency {
	return []Currency{
		Currency{
			Name:        "GBP",
			CountryName: "United Kingdom",
			Symbol:      "£",
		},
		Currency{
			Name:        "USD",
			CountryName: "United States",
			Symbol:      "$",
		},
		Currency{
			Name:        "EUR",
			CountryName: "Europe",
			Symbol:      "€",
		},
	}
}

// GetTaxes is a getter method to get the associated taxes in a Region
func (r *Region) GetTaxes() ([]Tax, error) {
	taxes := r.Taxes
	return taxes, nil
}

// JSON returns a JSON representation of the region
func (r *Region) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Region Object")
	}

	return b, nil
}
