package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	want := Product{
		ID:           3847132818,
		Name:         "Cup",
		Desc:         "A Nice Cup",
		Colour:       "White",
		SKU:          "C01-W",
		BasePrice:    5.99,
		BaseCurrency: "GBP",
		LocalPrice:   5.99,
	}

	if got := NewProduct("Cup", "A Nice Cup", "White", "C01-W", "GBP", 5.99); *got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestStoreProduct(t *testing.T) {
	// set up test data
	Seed()
	want := 4

	testProd := NewProduct("Cup", "A Nice Green Cup", "Green", "C02-G", "GBP", 5.99)

	err := StoreProduct(testProd)
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if len(KnownProducts.Products) != want {
		t.Errorf("got: %v, want: %v", len(KnownProducts.Products), want)
	}
}

func TestGenerateProductID(t *testing.T) {
	var want uint32 = 2313392227
	got := generateProductID("name", "Short Description", "colour", "sku")

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetProductOK(t *testing.T) {
	// initialise test products
	Seed()

	want := Product{
		ID:           2992948790,
		Name:         "Coffee Mug",
		Desc:         "A Nice Mug",
		Colour:       "White",
		SKU:          "CM01-W",
		BasePrice:    5.99,
		BaseCurrency: "GBP",
		LocalPrice:   5.99,
	}

	got, err := GetProduct("CM01-W")
	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if *got != want {
		t.Errorf("got: %v, want: %v", *got, want)
	}
}

func TestGetProductError(t *testing.T) {
	// initialise test products
	Seed()

	want := "No Product Matches SKU: Unknown"

	_, err := GetProduct("Unknown")
	if err == nil {
		t.Fatalf("Expected Error, got nil")
	}

	if err.Error() != want {
		t.Errorf("got: %v, want: %v", err.Error(), want)
	}
}

func TestGetProducts(t *testing.T) {
	// initialise test products
	Seed()
	want := seedProducts()

	got := GetProducts()

	if !reflect.DeepEqual(got, want.Products) {
		t.Errorf("got: %v, want: %v", got, want.Products)
	}
}

func TestProductJSON(t *testing.T) {
	// initialise test data
	Seed()

	want := Product{
		ID:           2992948790,
		Name:         "Coffee Mug",
		Desc:         "A Nice Mug",
		Colour:       "White",
		SKU:          "CM01-W",
		BasePrice:    5.99,
		BaseCurrency: "GBP",
		LocalPrice:   5.99,
	}

	testProduct, err := GetProduct("CM01-W")
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	bytes, err := testProduct.JSON()
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	var got Product
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestProductsJSON(t *testing.T) {
	// initialise test data
	Seed()

	want := seedProducts()

	testProductSlice := GetProducts()

	testProducts := Products{testProductSlice}

	bytes, err := testProducts.JSON()
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	var got Products
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
