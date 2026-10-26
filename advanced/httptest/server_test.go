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

// this test based on this post
// https://golang.cafe/blog/golang-httptest-example
//
// NOTE:
// This file just test simple success cases
// in practice, you should include more fail cases on your test file
//
// to run this file
// use go test -v .

// STEPS
// 1. define your handler
// 2. create a new [http.Request]
// 3. create a new [httptest.ResponseRecorder]
// 4. run handler(w, req)
// 5. get a resp by [httptest.ResponseRecorder.Result]
// 6. check if this response ok

// toUpper is a handler that parse word in query then write upper to response body
// Req: /upper?word=cat
// Res: CAT
func toUpper(w http.ResponseWriter, r *http.Request) {
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

func testUpperSuccess(input, output string) error {
	val := url.Values{
		"word": []string{input},
	}
	target := fmt.Sprintf("%s?%s", "/uppper", val.Encode())
	req := httptest.NewRequest(http.MethodGet, target, nil)

	// NOTE: NewRecorder implement interface [http.ResponseWriter]
	// We use handler (w, req) then get output from w
	w := httptest.NewRecorder()

	toUpper(w, req)

	// get a [http.Response] from Result function
	// the recorder will create a new Response based on the written data
	resp := w.Result()

	// then check if resp ok
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error, want: %d, got: %d", http.StatusOK, resp.StatusCode)
	}

	upperWord, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading body data: %s", err)
	}

	if string(upperWord) != output {
		return fmt.Errorf("error while comparing result, want: %s, got: %s", output, string(upperWord))
	}

	return nil
}

func TestServer(t *testing.T) {
	// create test cases
	testCases := [][]string{
		{"cat", "CAT"},
		{"bat", "BAT"},
		{"nix", "NIX"},
		{"golang", "GOLANG"},
		{"http", "HTTP"},
	}

	for _, c := range testCases {
		err := testUpperSuccess(c[0], c[1])
		if err != nil {
			t.Errorf("error while testing case: input -> %s, output -> %s, err: %s", c[0], c[1], err)
		}
	}
}
