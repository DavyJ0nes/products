package models

var (
	// KnownLocations is the in memory representation of the known Locations
	KnownLocations []Location
	// KnownCurrencies is the in memory representation of the known Currencies
	KnownCurrencies []Currency
	// KnownProducts is the in memory representation of the known Products
	KnownProducts Products
	// KnownTransactions is the in memory representation of the known Transactions
	KnownTransactions Transactions
)

// Seed initialises in memory data
func Seed() {
	KnownLocations = getKnownLocations()
	KnownCurrencies = getKnownCurrencies()
	KnownProducts = seedProducts()
	KnownTransactions = seedTransactions()
}

func seedProducts() Products {
	return Products{
		[]Product{
			{
				ID:           2992948790,
				Name:         "Coffee Mug",
				Desc:         "A Nice Mug",
				Colour:       "White",
				SKU:          "CM01-W",
				BasePrice:    5.99,
				BaseCurrency: "GBP",
				LocalPrice:   5.99,
			},
			{
				ID:           2992948790,
				Name:         "Coaster",
				Desc:         "Cork Coaster",
				Colour:       "Brown",
				SKU:          "Co01-B",
				BasePrice:    2.50,
				BaseCurrency: "USD",
				LocalPrice:   2.50,
			},
			{
				ID:           2992948790,
				Name:         "Glass Tumbler",
				Desc:         "Whiskey Glass",
				Colour:       "Glass",
				SKU:          "GT01-G",
				BasePrice:    12.99,
				BaseCurrency: "EUR",
				LocalPrice:   12.99,
			},
		},
	}

}

// TODO (davy): Once data store is added this will need to be abstracted
func getKnownLocations() []Location {
	uk := Location{
		Name: "United Kingdom",
		Currency: Currency{
			Name:        "GBP",
			CountryName: "United Kingdom",
			Symbol:      "£",
		},
		Taxes: []Tax{{"VAT", 0.2, 0.0}},
	}

	pasadena := Location{
		Name: "Pasadena, CA, USA",
		Currency: Currency{
			Name:        "USD",
			CountryName: "United States",
			Symbol:      "$",
		},
		Taxes: []Tax{
			{"Sales Tax", 0.095, 0.0},
			{"Federal Tax", 0.095, 0.0},
		},
	}

	fra := Location{
		Name: "France",
		Currency: Currency{
			Name:        "EUR",
			CountryName: "EU Zone",
			Symbol:      "€",
		},
		Taxes: []Tax{{"VAT", 0.2, 0.0}},
	}

	ger := Location{
		Name: "Germany",
		Currency: Currency{
			Name:        "EUR",
			CountryName: "EU Zone",
			Symbol:      "€",
		},
		Taxes: []Tax{{"VAT", 0.19, 0.0}},
	}

	return []Location{uk, pasadena, fra, ger}
}

// getKnownCurrencies looks for already defined currencies
// currently currencies are stored statically, this will need to be refactored once database has been added
func getKnownCurrencies() []Currency {
	return []Currency{
		{
			Name:        "GBP",
			CountryName: "United Kingdom",
			Symbol:      "£",
		},
		{
			Name:        "USD",
			CountryName: "United States",
			Symbol:      "$",
		},
		{
			Name:        "EUR",
			CountryName: "Europe",
			Symbol:      "€",
		},
	}
}

// TODO (davy): To be implemented
func seedTransactions() Transactions {
	return []Transaction{}
}
