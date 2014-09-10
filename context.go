package gofamilysearch

import "net/http"

type Context struct {
	environment *Environment
	accessToken string
	client      *http.Client
}

// Create a Context. The Context is specific to a user
// pass in client to allow running on appengine
func NewContext(environment *Environment, accessToken string, client *http.Client) *Context {
	return &Context{
		environment: environment,
		accessToken: accessToken,
		client:      client,
	}
}
