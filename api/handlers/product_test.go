package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davyj0nes/products/api/models"
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
			want:         `{"id":2992948790,"name":"Coaster","desc":"Cork Coaster","colour":"Brown","sku":"Co01-B","price":250,"base_currency":"USD","local_price":250}`,
		},
		{
			name:         "product doesn't exist",
			url:          "/api/v1/product/unknown",
			sku:          "unknown",
			responseCode: http.StatusNotFound,
			want:         "Error: No Product Matches SKU: unknown",
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

func TestNewProduct(t *testing.T) {
	// seeding in memory data
	models.Seed()

	testCases := []struct {
		name         string
		input        string
		responseCode int
		want         string
	}{
		{
			name:         "empty input",
			input:        ``,
			responseCode: http.StatusBadRequest,
			want:         "received no data",
		},
		{
			name:         "new product",
			input:        `{"name": "Wired Mouse","description": "Microsoft Wired Mouse","colour": "Black","sku":"Mo01-B","currency":"USD","price":2495}`,
			responseCode: http.StatusOK,
			want:         `{"id":2050674932,"name":"Wired Mouse"}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/api/v1/product", bytes.NewBuffer([]byte(tt.input)))
			if err != nil {
				t.Errorf("Unexpected Error: %v", err.Error())
			}

			newProduct(w, req)
			resp := w.Result()

			checkResponseCode(t, tt.responseCode, resp.StatusCode)

			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if tt.responseCode != http.StatusOK {
				var errMsg Error
				err = json.Unmarshal(body, &errMsg)
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}

				if errMsg.Message != tt.want {
					t.Errorf("got: %v, want: %v", errMsg.Message, tt.want)
				}
			} else {
				if string(body) != tt.want+"\n" {
					t.Errorf("got:  %v, want: %v", string(body), tt.want)
				}
			}
		})
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
