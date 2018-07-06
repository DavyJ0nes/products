package models

// APIElement describes the behaviour of an element of the API
// Elements that implement this interface allow for easier pulling of their data
type APIElement interface {
	JSON() ([]byte, error)
}

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
			},
			{
				ID:           2992948790,
				Name:         "Coaster",
				Desc:         "Cork Coaster",
				Colour:       "Brown",
				SKU:          "Co01-B",
				BasePrice:    2.50,
				BaseCurrency: "USD",
			},
			{
				ID:           2992948790,
				Name:         "Glass Tumbler",
				Desc:         "Whiskey Glass",
				Colour:       "Glass",
				SKU:          "GT01-G",
				BasePrice:    12.99,
				BaseCurrency: "EUR",
			},
		},
	}

}

func seedTransactions() Transactions {
	return []Transaction{}
}
