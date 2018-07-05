package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNonExistantProduct(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/product/45", nil)
	if err != nil {
		t.Errorf("Unexpected Error: %v", err.Error())
	}

	getProduct(w, req)
	resp := w.Result()

	checkResponseCode(t, http.StatusNotFound, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var errMsg Error
	json.Unmarshal(body, &errMsg)
	if errMsg.Message != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", errMsg.Message)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
