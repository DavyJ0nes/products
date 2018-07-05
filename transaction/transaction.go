package transaction

import (
	"github.com/davyj0nes/products-api/product"
	"time"
)

type Transaction struct {
	ID         string
	Datetime   time.Time
	RegionName string
	Products   []product.Product
	Subtotal   float64
	TaxTotal   float64
	Total      float64
}

func exampleFunction(s string) int {
	return len(s)
}
