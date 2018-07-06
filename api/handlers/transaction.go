package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davyj0nes/products/api/models"
	log "github.com/sirupsen/logrus"
)

// NewTransactionInput is used for the JSON POST input to the NewTransaction
type NewTransactionInput struct {
	Location    string   `json:"location,omitempty"`
	ProductSKUS []string `json:"product_skus,omitempty"`
}

// TranProductOutput is used to generate JSON response for all products
type TranProductOutput struct {
	ProductQuantity int     `json:"product_quantity,omitempty"`
	ProductName     string  `json:"product_name,omitempty"`
	Price           float64 `json:"price,omitempty"`
}

// TransactionOutput is used for the JSON response for a transaction
type TransactionOutput struct {
	OrderID           string              `json:"order_id,omitempty"`
	FormattedProducts []TranProductOutput `json:"formatted_products,omitempty"`
	FormattedDateTime string              `json:"formatted_date_time,omitempty"`
	Subtotal          float64             `json:"subtotal,omitempty"`
	Taxtotal          float64             `json:"taxtotal,omitempty"`
	Total             float64             `json:"total,omitempty"`
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	log.Info("Received Request: ", "newTransaction")
	var (
		tranInput  NewTransactionInput
		tranOutput TransactionOutput
		err        error
	)

	b, err := ioutil.ReadAll(r.Body)
	checkError(w, err)
	err = json.Unmarshal(b, &tranInput)
	checkError(w, err)

	tran, err := models.NewTransaction(tranInput.Location)
	log.Infof("Created New Transaction: %s", tran.ID)

	var products []models.Product
	for _, sku := range tranInput.ProductSKUS {
		p, err := models.GetProduct(sku)
		if err != nil {
			checkError(w, err)
			return
		}
		products = append(products, *p)
	}

	tran.AddProducts(products)

	err = tran.CalcSubtotal()
	if err != nil {
		checkError(w, err)
		return
	}

	tran.CalcTaxTotal()

	tran.CalcTransactionTotal()

	// store completed transaction in data store
	err = models.StoreTransaction(tran)
	if err != nil {
		checkError(w, err)
		return
	}

	log.Infof("Finished Calculating Transaction: %v", *tran)
	// Set up output
	tranOutput.OrderID = tran.ID
	tranOutput.FormattedProducts = formatProductOutput(tran.Products)
	tranOutput.FormattedDateTime = tran.Datetime.Format("02-01-2006 15:04:05")
	tranOutput.Subtotal = tran.Subtotal
	tranOutput.Taxtotal = tran.TaxTotal
	tranOutput.Total = tran.Total

	generateJSONResponse(w, http.StatusOK, tranOutput)
}

func allTransactions(w http.ResponseWriter, r *http.Request) {
}

func getTransaction(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// vars["id"] // id of the transaction
}

// formatProductOutput translates product information into one that is better for the transaction output
// TODO (davy): Product Quantity need to be improved
func formatProductOutput(products []models.Product) []TranProductOutput {
	var output []TranProductOutput
	for _, p := range products {
		fp := TranProductOutput{
			Price:           p.LocalPrice,
			ProductName:     p.Name,
			ProductQuantity: 1,
		}
		output = append(output, fp)
	}
	return output
}
