package gofamilysearch

// Relationship contains information about a couple or a parent-child relationship
type Relationship struct {
	ID      string           `json:"id"`
	Type    string           `json:"type"`
	Links   map[string]*Link `json:"links"`
	Facts   []*Fact          `json:"facts"`
	Person1 ResourceRef      `json:"person1"`
	Person2 ResourceRef      `json:"person2"`
}
