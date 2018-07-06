package models

import "testing"

func TestNewProduct(t *testing.T) {
	want := Product{
		ID:           2992948790,
		Name:         "Cup",
		Desc:         "A Nice Cup",
		BasePrice:    5.99,
		BaseCurrency: "GBP",
	}

	if got := NewProduct("Cup", "A Nice Cup", "GBP", 5.99); *got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGenerateProductID(t *testing.T) {
	var want uint32 = 3581725991
	got := generateProductID("name", "Short Description", 9.99)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
