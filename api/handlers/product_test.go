package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davyj0nes/products/api/v1/models"
	"github.com/gorilla/mux"
)

func TestGetProduct(t *testing.T) {
	// seeding in memory data
	models.Seed()

	testCases := []struct {
		name         string
		url          string
		sku          string
		responseCode int
		want         string
	}{
		{
			name:         "product exists",
			url:          "/api/v1/product/Co01-B",
			sku:          "Co01-B",
			responseCode: http.StatusOK,
			want:         `{"id":2992948790,"name":"Coaster","desc":"Cork Coaster","colour":"Brown","sku":"Co01-B","price":2.5,"base_currency":"USD"}`,
		},
		{
			name:         "product doesn't exist",
			url:          "/api/v1/product/unknown",
			sku:          "unknown",
			responseCode: http.StatusNotFound,
			want:         "Problem finding product: No Product Matches SKU",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Errorf("Unexpected Error: %v", err.Error())
			}

			// Having to do this as a bit of a hack to get URL Params to work
			vars := map[string]string{
				"sku": tt.sku,
			}

			req = mux.SetURLVars(req, vars)

			getProduct(w, req)
			resp := w.Result()

			checkResponseCode(t, tt.responseCode, resp.StatusCode)

			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if tt.responseCode == http.StatusNotFound {
				var errMsg Error
				err = json.Unmarshal(body, &errMsg)
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}

				if errMsg.Message != tt.want {
					t.Errorf("got: %v, want: %v", errMsg.Message, tt.want)
				}
			} else {
				// body string has new line at the end so need to add this here
				// TODO (davy): Fix this
				if string(body) != tt.want+"\n" {
					t.Errorf("\ngot:  %v\nwant: %v", string(body), tt.want)
				}
			}
		})
	}
}

func TestAllProducts(t *testing.T) {
	// seeding in memory data
	models.Seed()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/product/all", nil)
	if err != nil {
		t.Errorf("Unexpected Error: %v", err.Error())
	}

	allProducts(w, req)
	resp := w.Result()

	checkResponseCode(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var productJSON models.Products
	json.Unmarshal(body, &productJSON)
	if len(productJSON.Products) != 3 {
		t.Errorf("got: %v, want: %v", len(productJSON.Products), 3)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
