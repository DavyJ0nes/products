package models

import (
	"testing"
)

func TestCalcLocalPrice(t *testing.T) {
	rate := 1.128989
	testCases := []struct {
		testName     string
		productPrice float64
		want         float64
	}{
		{"Basic Input", 5.99, 6.77},
		{"No Price", 0.0, 0.0},
		{"Odd Price", 1.23, 1.39},
		{"Negative Price", -1.23, -1.39},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			if got := calcLocalPrice(tt.productPrice, rate); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetLocalRate(t *testing.T) {
	_, err := getLocalRate("USD", "GBP")
	if err != nil {
		t.Errorf("Unexpected Error: %v", err.Error())
	}
}

func TestGenerateTransactionID(t *testing.T) {
	id, err := generateTransactionID()
	if err != nil {
		t.Errorf("Unexpected Error: %v", err.Error())
	}

	if len(id) != 36 {
		t.Errorf("got: %d, want: %d", len(id), 36)
	}
}
