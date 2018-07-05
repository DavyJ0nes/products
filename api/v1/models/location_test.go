package models

import (
	"reflect"
	"testing"
)

func TestNewLocationOK(t *testing.T) {
	inputTaxes := []Tax{
		{
			Name:   "VAT",
			Amount: 0.2, // 20%
		},
	}

	want := Location{
		Name:     "UK",
		Currency: Currency{"GBP", "United Kingdom", "£"},
		Taxes:    inputTaxes,
	}

	got, err := NewLocation("UK", "GBP", inputTaxes)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestNewLocationError(t *testing.T) {
	want := "Issue Validating Currency: Currency Not Found: WAK"

	_, err := NewLocation("Wakana", "WAK", []Tax{})
	if err == nil {
		t.Errorf("Expected Error but got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}
