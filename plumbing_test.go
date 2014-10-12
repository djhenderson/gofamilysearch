package gofamilysearch

import (
	"testing"
)

func TestReadDiscoveryResource(t *testing.T) {
	testSetup()
	defer testTeardown()

	testRespond(t, "GET", "/platform/collections/tree", nil, "platform_collections_tree.json")

	want := map[string]string{
		"person": "https://sandbox.familysearch.org/platform/tree/persons/{pid}",
		"person-with-relationships": "https://sandbox.familysearch.org/platform/tree/persons-with-relationships",
		"persons": "https://sandbox.familysearch.org/platform/tree/persons",
	}
	templates, err := testClient().readDiscoveryResource(testServer.URL)
	if err != nil {
		t.Error(err)
	}
	for k, v := range want {
		if templates[k] != v {
			t.Errorf("readTemplates returned %s, want %s", templates[k], v)
		}
	}
}
