package http_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var successCases = [][]string{
	{"cat", "CAT"},
	{"bat", "BAT"},
	{"nix", "NIX"},
	{"golang", "GOLANG"},
	{"http", "HTTP"},
}

// upperHandler convert word on url query into upper the write to body
func upperHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request")
		return
	}

	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, strings.ToUpper(word))
}

// this test file based on post https://golang.cafe/blog/golang-httptest-example

// It introduce two method to test client
// for golang latter 1.27.3, in-memory server test is highly recommend

type TestClient struct {
	httpClient *http.Client
	url        string
}

func NewClient(c *http.Client, url string) *TestClient {
	return &TestClient{httpClient: c, url: url}
}

// UpperCase client send a GET Request for getting the upper of a word
func (c TestClient) UpperCase(word string) (string, error) {
	// for a loopback server test, you must config the url
	res, err := c.httpClient.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", fmt.Errorf("Unable to complete Get request: %s", err)
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Unable to read response data: %s", err)
	}

	return string(out), nil
}

// STEPS
// 1. create a new test server (loopback)
// 2. create a client based on this server URL
// 3. send request to this server
// 4. check result
// 5. close server
func TestLoopBackClient(t *testing.T) {
	// NOTE: [httptest.NewServer] use
	// local network loopback interface.
	// which cost more resource then in-memory Server
	//
	// If your golang version is latter then 1.27.3, you should use [httptest.NewTestServer] to create an in-memory server instead
	srv := httptest.NewServer(http.HandlerFunc(upperHandler))

	// remeber to close this server when test end
	defer srv.Close()

	// this is a loopback server, so it will not be empty
	t.Logf("creating a new client with url: %s", srv.URL)
	c := NewClient(srv.Client(), srv.URL)

	for _, ca := range successCases {
		res, err := c.UpperCase(ca[0])
		if err != nil {
			t.Errorf("error while calling UpperCase, err: %s", err)
		}
		if res != ca[1] {
			t.Errorf("expected: %s, got: %s", ca[1], res)
		}
	}
}

// This method require golang latter 1.27.3
// STEPS
// 1. create a new in-memory server
// 2. create a client based on [httptest.Server.Client](), the url can be random
// 3. send request
// 4. check result
// 5. close test server
//
// func TestInMomoryClient(t *testing.T) {
// 	// NOTE: This in-memory server just support after go 1.27.3
//
// 	// Highly recommend this in-memory test if you have a lot of servers to created on test
// 	// or your test cost a lot of time on loopback server
// 	srv := httptest.NewTestServer(t, http.HandlerFunc(upperHandler))
//
// 	defer srv.Close()
//
// 	// WARN: It will be empty
// 	// so you should use Client function to get a client
// 	t.Logf("creating a new client with url: %s", srv.URL)
//
// 	// just config a random url bcs in-memory server client will send all request to test server
// 	// NOTE: Regradless the url, all request will be sent to test server
// 	c := NewClient(srv.Client(), "https://go.dev")
//
// 	for _, ca := range successCases {
// 		res, err := c.UpperCase(ca[0])
// 		if err != nil {
// 			t.Errorf("error while calling UpperCase, err: %s", err)
// 		}
// 		if res != ca[1] {
// 			t.Errorf("expected: %s, got: %s", ca[1], res)
// 		}
// 	}
// }
