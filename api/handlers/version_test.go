package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	// TESTING /version Status Code
	w := httptest.NewRecorder()

	testBuildTime := time.Now().Format("2006-01-02_03:04:05")
	testCommit := "abcd123"
	testRelease := "0.0.1"

	h := version(testBuildTime, testCommit, testRelease)
	h(w, nil)
	resp := w.Result()

	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d", have, want)
	}

	// TESTING /version Response Body
	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		BuildTime string `json:"build_time"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}

	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}

	if info.Release != testRelease {
		t.Errorf("Release version is wrong. Have: %s, want: %s", info.Release, testRelease)
	}
	if info.BuildTime != testBuildTime {
		t.Errorf("Build time is wrong. Have: %s, want: %s", info.BuildTime, testBuildTime)
	}
	if info.Commit != testCommit {
		t.Errorf("Commit is wrong. Have: %s, want: %s", info.Commit, testCommit)
	}
}
