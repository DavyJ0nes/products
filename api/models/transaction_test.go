package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// This tests the full transaction process with a few different inputs
// TODO (davy): mock conversion service
func TestFullTransaction(t *testing.T) {
	// set up test data
	Seed()

	allProducts := seedProducts().Products
	product1, _ := GetProduct("CM01-W")

	testCases := []struct {
		testName string
		location string
		products []Product
		want     float64
	}{
		{
			"One Product UK Test",
			"United Kingdom",
			[]Product{*product1},
			7.19,
		},
		{
			"Three Products UK Test",
			"United Kingdom",
			allProducts,
			23.32,
		},
		{
			"One Product France Test",
			"France",
			[]Product{*product1},
			8.12,
		},
		{
			"Three Products France Test",
			"France",
			allProducts,
			26.26,
		},
		{
			"One Product Pasadena, CA, USA Test",
			"Pasadena, CA, USA",
			[]Product{*product1},
			8.73,
		},
		{
			"Three Products Pasadena, CA, USA Test",
			"Pasadena, CA, USA",
			allProducts,
			28.11,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			// Create New Transaction in Desired Location
			testTran, err := NewTransaction(tt.location)
			if err != nil {
				t.Fatalf("Unexpected Error: %v", err)
			}

			// Add Products to Transaction
			testTran.AddProducts(tt.products)

			// Calculate Total of Product Prices
			err = testTran.CalcSubtotal()
			if err != nil {
				t.Fatalf("Unexpected Error: %v", err)
			}
			fmt.Println("Subtotal:", testTran.Subtotal)

			// Calculate Total Tax
			testTran.CalcTaxTotal()
			fmt.Println("TaxTotal:", testTran.TaxTotal)

			// Calculate Final Total
			testTran.CalcTransactionTotal()
			fmt.Println("FinalTotal:", testTran.Total)
			fmt.Println("ConversionRate:", testTran.ConversionRate)

			if got := testTran.Total; got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
func TestNewTransaction(t *testing.T) {
	want, err := GetLocation("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	got, err := NewTransaction("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	// Only testing that function returns the correct location as need to
	// figure out how to test dates and uuid generation
	if !reflect.DeepEqual(*got.Location, *want) {
		t.Errorf("got: %v, want: %v", *got.Location, *want)
	}
}

func TestNewTransactionError(t *testing.T) {
	want := "Problem Getting Location Information: Location Not Found: Unknown"

	_, err := NewTransaction("Unknown")
	if err == nil {
		t.Fatalf("Expected Error, got: nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}

func TestAddProducts(t *testing.T) {
	want := 2

	testProducts := []Product{
		{
			ID:           3847132818,
			Name:         "Cup",
			Desc:         "A Nice Cup",
			Colour:       "White",
			SKU:          "C01-W",
			BasePrice:    5.99,
			BaseCurrency: "GBP",
		},
		{
			ID:           3847132818,
			Name:         "Cup",
			Desc:         "A Nice Cup",
			Colour:       "White",
			SKU:          "C01-W",
			BasePrice:    5.99,
			BaseCurrency: "GBP",
		},
	}

	testTransaction, err := NewTransaction("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	testTransaction.AddProducts(testProducts)

	if len(testTransaction.Products) != want {
		t.Errorf("got: %v, want: %v", len(testTransaction.Products), want)
	}
}

// TODO (davy): mock conversion service
func TestCalcTotals(t *testing.T) {
	wantSubtotal := 11.98
	wantTaxTotal := 2.40
	wantTotal := 14.38

	testProducts := []Product{
		{
			ID:           3847132818,
			Name:         "Cup",
			Desc:         "A Nice Cup",
			Colour:       "White",
			SKU:          "C01-W",
			BasePrice:    5.99,
			BaseCurrency: "GBP",
		},
		{
			ID:           3847132818,
			Name:         "Cup",
			Desc:         "A Nice Cup",
			Colour:       "White",
			SKU:          "C01-W",
			BasePrice:    5.99,
			BaseCurrency: "GBP",
		},
	}

	testTransaction, err := NewTransaction("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	testTransaction.AddProducts(testProducts)

	err = testTransaction.CalcSubtotal()
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if testTransaction.Subtotal != wantSubtotal {
		t.Errorf("got: %v, want: %v", testTransaction.Subtotal, wantSubtotal)
	}

	testTransaction.CalcTaxTotal()
	if testTransaction.TaxTotal != wantTaxTotal {
		t.Errorf("got: %v, want: %v", testTransaction.TaxTotal, wantTaxTotal)
	}

	testTransaction.CalcTransactionTotal()
	if testTransaction.Total != wantTotal {
		t.Errorf("got: %v, want: %v", testTransaction.Total, wantTotal)
	}
}

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

// TODO (davy): mock conversion service
func TestGetLocalRate(t *testing.T) {
	_, err := getLocalRate("USD", "GBP")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err.Error())
	}
}

func TestGenerateTransactionID(t *testing.T) {
	id, err := generateTransactionID()
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err.Error())
	}

	if len(id) != 36 {
		t.Errorf("got: %d, want: %d", len(id), 36)
	}
}

func TestTransactionJSON(t *testing.T) {
	want, err := NewTransaction("United Kingdom")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	bytes, err := want.JSON()
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	var got Transaction
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	// Due to the location being a pointer
	// When Unmarshalling it creates a copy and therefore the Location pointer of got is a different memory location
	// So for ease am just testing the uuid
	if got.ID != want.ID {
		t.Errorf("got: %v, want: %v", got.ID, want.ID)
	}
}
