package models

import "testing"

func TestNewProduct(t *testing.T) {
	want := Product{
		ID:           3847132818,
		Name:         "Cup",
		Desc:         "A Nice Cup",
		Colour:       "White",
		SKU:          "C01-W",
		BasePrice:    5.99,
		BaseCurrency: "GBP",
	}

	if got := NewProduct("Cup", "A Nice Cup", "White", "C01-W", "GBP", 5.99); *got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGenerateProductID(t *testing.T) {
	var want uint32 = 2313392227
	got := generateProductID("name", "Short Description", "colour", "sku")

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
