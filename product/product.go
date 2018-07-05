package product

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Regions     []Region
}

type Region struct {
	Name     string
	Currency string
	Taxes    []Tax
}

type Tax struct {
	Name   string
	Amount float64
}

type Currency struct {
	Name        string
	CountryName string
	Symbol      rune
}

type Currencies []Currency

func exampleFunction(s string) int {
	return len(s)
}
