package fstesting

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	Mux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses.
	Server *httptest.Server
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func Setup() {
	// test server
	Mux = http.NewServeMux()
	Server = httptest.NewServer(Mux)
}

// teardown closes the test HTTP server.
func Teardown() {
	Server.Close()
}

func TestMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

func TestHeader(t *testing.T, r *http.Request, header string, want string) {
	if value := r.Header.Get(header); want != value {
		t.Errorf("Header %s = %s, want: %s", header, value, want)
	}
}
