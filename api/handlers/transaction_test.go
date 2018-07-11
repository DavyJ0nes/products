package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davyj0nes/products/api/models"
)

func TestNewTransaction(t *testing.T) {
	// seeding in memory data
	models.Seed()

	// start mock converstion rate API server
	server := mockConversionServer()
	defer server.Close()

	// overwriting global config variable
	conversionAPIURL = server.URL

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
			name:         "new transaction",
			input:        `{"location": "United Kingdom","product_skus": ["CM01-W","Co01-B","GT01-G"]}`,
			responseCode: http.StatusOK,
			// only checking total price here.
			// TODO (davy): Will need to update this in future.
			want: "34.03",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/api/v1/transaction", bytes.NewBuffer([]byte(tt.input)))
			if err != nil {
				t.Errorf("Unexpected Error: %v", err.Error())
			}

			newTransaction(w, req)
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
				var tempTransaction models.Transaction
				json.Unmarshal(body, &tempTransaction)
				totalString := fmt.Sprintf("%.2f", tempTransaction.Total)
				if totalString != tt.want {
					t.Errorf("\ngot:  %v\nwant: %v", totalString, tt.want)
				}
			}
		})
	}
}

func mockConversionServer() *httptest.Server {
	f := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"GBP_USD":{"val":1.321406}}`)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
