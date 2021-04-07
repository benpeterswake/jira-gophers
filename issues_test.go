package jira

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestClient_Search(t *testing.T) {
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

	query := "test"
	expected := searchResult{
		Issues: []Issue{
			{
				Expand:         "test",
				ID:             "test",
				Self:           "test",
				Key:            "test",
				Fields:         &IssueFields{},
				RenderedFields: &IssueRenderedFields{},
				Changelog:      &Changelog{},
				Transitions:    []Transition{},
				Names:          map[string]string{},
			},
		},
	}

	testMux.HandleFunc("/rest/api/3/search", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		jql := params.Get("jql")
		if jql != query {
			t.Fatal("Expected", query, "but got", jql)
		}
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(&expected)
		if err != nil {
			t.Fatal(err)
		}
		return
	})

	actual, err := testClient.Search("test", nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(actual) != len(expected.Issues) {
		t.Fatal("Expected matching issues lists")
	}

	if actual[0].ID != expected.Issues[0].ID {
		t.Fatal("Expected ID", expected.Issues[0].ID, "but got", actual[0].ID)
	}

	if actual[0].Key != expected.Issues[0].Key {
		t.Fatal("Expected Key", expected.Issues[0].Key, "but got", actual[0].Key)
	}

	if actual[0].Self != expected.Issues[0].Self {
		t.Fatal("Expected Self", expected.Issues[0].Self, "but got", actual[0].Self)
	}
}
