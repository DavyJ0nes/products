package models

import (
	"reflect"
	"testing"
)

func TestNewRegionOK(t *testing.T) {
	inputTaxes := []Tax{
		{
			Name:   "VAT",
			Amount: 0.2, // 20%
		},
	}

	want := Region{
		Name:         "UK",
		CurrencyName: "GBP",
		Taxes:        inputTaxes,
	}

	got, err := NewRegion("UK", "GBP", inputTaxes)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestNewRegionError(t *testing.T) {
	want := "Issue Validating Currency: Currency Not Found: WAK"

	_, err := NewRegion("Wakana", "WAK", []Tax{})
	if err == nil {
		t.Errorf("Expected Error but got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}
