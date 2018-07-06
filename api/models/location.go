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

// Location is a specific geographic area with specific tax laws
// Generally will refer to the location that a transaction is carried out in
type Location struct {
	Name     string
	Currency Currency
	Taxes    []Tax
}

// NewLocation is a factory for creating new regions unsurprisingly
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

func getKnownLocations() []Location {
	uk := Location{
		Name: "United Kingdom",
		Currency: Currency{
			Name:        "GBP",
			CountryName: "United Kingdom",
			Symbol:      "£",
		},
		Taxes: []Tax{{"VAT", 0.2}},
	}

	pasadena := Location{
		Name: "Pasadena, CA, USA",
		Currency: Currency{
			Name:        "USD",
			CountryName: "United States",
			Symbol:      "$",
		},
		Taxes: []Tax{
			{"Sales Tax", 0.095},
		},
	}

	fra := Location{
		Name: "France",
		Currency: Currency{
			Name:        "EUR",
			CountryName: "EU Zone",
			Symbol:      "€",
		},
		Taxes: []Tax{{"VAT", 0.2}},
	}

	ger := Location{
		Name: "Germany",
		Currency: Currency{
			Name:        "EUR",
			CountryName: "EU Zone",
			Symbol:      "€",
		},
		Taxes: []Tax{{"VAT", 0.19}},
	}

	return []Location{uk, pasadena, fra, ger}
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
func (l *Location) GetTaxes() ([]Tax, error) {
	taxes := l.Taxes
	return taxes, nil
}

// JSON returns a JSON representation of the region
func (l *Location) JSON() ([]byte, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Problem Encoding Region Object")
	}

	return b, nil
}
