package jira

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the Jira client being tested.
	testClient Client

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a jira.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	addr := strings.ReplaceAll(testServer.URL, "http://", "")
	testClient = NewClient(addr, addr, "http", "test", "test", "test")
}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}
