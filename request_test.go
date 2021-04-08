package jira

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSendRequest(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		resp := OAuthResponse{}
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			t.Fatal(err)
		}

		return
	})

	testMux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		return
	})

	method := http.MethodGet
	url := testClient.getScheme() + "://" + testClient.getBaseURL() + "/issues"
	req, err := testClient.newRequest(method, url, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp, err := testClient.sendRequest(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestNewRequest(t *testing.T) {
	method := http.MethodGet
	url := "https://example.com"

	req, err := testClient.newRequest(method, url, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	if req.Method != method {
		t.Fatal("Wanted request method", method, "but got", req.Method)
	}

	if req.URL.String() != url {
		t.Fatal("Wanted request url", url, "but got", req.URL.String())
	}

	if req.Header.Get("Content-Type") != "application/json" {
		t.Fatal("Wanted header content type to be application/json")
	}

	if req.Header.Get("Accept") != "application/json" {
		t.Fatal("Wanted header accept to be application/json")
	}
}
