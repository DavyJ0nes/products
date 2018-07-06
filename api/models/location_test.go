package models

import (
	"encoding/json"
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
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestNewLocationError(t *testing.T) {
	want := "Issue Validating Currency: Currency Not Found: WAK"

	_, err := NewLocation("Wakanda", "WAK", []Tax{})
	if err == nil {
		t.Fatalf("Expected Error but got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}

func TestGetLocationOK(t *testing.T) {
	want := &Location{
		Name:     "United Kingdom",
		Currency: Currency{"GBP", "United Kingdom", "£"},
		Taxes: []Tax{
			{
				Name:   "VAT",
				Amount: 0.2, // 20%
			},
		},
	}

	got, err := GetLocation("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetLocationError(t *testing.T) {
	want := "Location Not Found: Unknown"

	_, err := GetLocation("Unknown")
	if err == nil {
		t.Fatalf("Expecting method to error, got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}
func TestGetTaxesOK(t *testing.T) {
	want := []Tax{
		{
			Name:   "VAT",
			Amount: 0.2, // 20%
		},
		{
			Name:   "City Tax",
			Amount: 0.05, // 5%
		},
	}

	testLocation, err := NewLocation("UK", "GBP", want)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	got, err := testLocation.GetTaxes()
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetTaxesEmpty(t *testing.T) {
	want := "No Taxes associated with Location"

	testLocation, err := NewLocation("UK", "GBP", []Tax{})
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	_, err = testLocation.GetTaxes()
	if err == nil {
		t.Fatalf("Expecting method to error, got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}

func TestLocationJSON(t *testing.T) {
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

	testLocation, err := NewLocation("UK", "GBP", inputTaxes)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	bytes, err := testLocation.JSON()
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	var got Location
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
