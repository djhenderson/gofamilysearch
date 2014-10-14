package gofamilysearch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"regexp"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func testSetup() {
	// test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
}

// teardown closes the test HTTP server.
func testTeardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

func testQueryParameters(t *testing.T, r *http.Request, want map[string]string) {
	r.ParseForm()
	eq := true
	if want == nil {
		eq = len(r.Form) == 0
	} else if len(r.Form) != len(want) {
		eq = false
	} else {
		for k, v := range want {
			if r.FormValue(k) != v {
				eq = false
				break
			}
		}
	}
	if !eq {
		t.Errorf("Query Parameters = %+v, want %+v", r.Form, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if value := r.Header.Get(header); want != value {
		t.Errorf("Header %s = %s, want: %s", header, value, want)
	}
}

var charsToTranslate = regexp.MustCompile("[./]")

func testRespond(t *testing.T, method string, url string, queryParams map[string]string, filename string) {
	if filename == "" {
		filename = charsToTranslate.ReplaceAllString(url[1:], "_")+".json"
	}
	testMux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, method)
		testQueryParameters(t, r, queryParams)
		contents, err := readFile("test_responses/" + filename)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(contents))
	})
}

func testClient() *Client {
	contents, err := readFile("test_responses/_well-known_app-meta.json")
	if err != nil {
		panic(err)
	}
	response := &discoveryResponse{}
	if err := json.Unmarshal(contents, response); err != nil {
		panic(err)
	}

	templates, err := generateTemplates(testServer.URL, response)
	if err != nil {
		panic(err)
	}

	// need to replace all hosts with our testServer.URL
	for k, v := range templates {
		v = strings.Replace(v, "https://sandbox.familysearch.org", testServer.URL, 1)
		v = strings.Replace(v, "https://integration.familysearch.org", testServer.URL, 1)
		templates[k] = v
	}

	return &Client{
		AccessToken: "accessToken",
		HTTPClient:  &http.Client{},
		Context: &Context{
			templates: templates,
		},
	}

}
