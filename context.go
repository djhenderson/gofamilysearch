package gofamilysearch

import "sync"

// The Context can be shared among go-routines
type Context struct {
	Environment string
	once        sync.Once
	templates   map[string]string
}
