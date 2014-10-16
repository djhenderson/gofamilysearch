package gofamilysearch

import "net/http"

// Client is specific to a user
// pass in http-client to allow running on appengine
type Client struct {
	Context     *Context
	AccessToken string
	Transport   http.RoundTripper
}
