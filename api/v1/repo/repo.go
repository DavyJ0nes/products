package repo

import "github.com/davyj0nes/products/api/v1/models"

var (
	currentID    uint32
	products     models.Products
	transactions models.Transactions
)

// create seed data
func init() {
	CreateProduct(models.Product{
		Name:  "Pen",
		Desc:  "Lamy Fountain Pen",
		Price: 5.99,
	})

	CreateProduct(models.Product{
		Name:  "Banana",
		Desc:  "Banana",
		Price: 0.60,
	})

	CreateProduct(models.Product{
		Name:  "Headphones",
		Desc:  "Sennheiser Headphones",
		Price: 35.0,
	})
}

// CreateProduct Adds New Product to in memory storage
func CreateProduct(p models.Product) {
	currentID++
	p.ID = currentID
	products = append(products, p)
}
