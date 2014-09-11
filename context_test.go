package gofamilysearch

import (
	"fmt"
	"github.com/rootsdev/gofamilysearch/fstesting"
	"net/http"
	"reflect"
	"testing"
)

func TestReadTemplates(t *testing.T) {
	fstesting.Setup()
	defer fstesting.Teardown()

	fstesting.Mux.HandleFunc("/.well-known/app-meta", func(w http.ResponseWriter, r *http.Request) {
		fstesting.TestMethod(t, r, "GET")
		fmt.Fprintf(w, `
<?xml version="1.0" ?>
<atom:feed xmlns:atom="http://www.w3.org/2005/Atom" xmlns:fs="http://familysearch.org/v1/" xmlns:gx="http://gedcomx.org/v1/">
  <atom:link rel="agent-template" template="https://sandbox.familysearch.org/platform/users/agents/{uid}{?access_token}" title="Agent" type="application/json,application/x-fs-v1+json,application/x-fs-v1+xml,application/x-gedcomx-v1+json,application/x-gedcomx-v1+xml,application/xml,text/html" accept="*/*" allow="GET"></atom:link>
  <atom:link rel="child-and-parents-relationships" href="/platform/tree/child-and-parents-relationships"></atom:link>
  <atom:link rel="child-relationships-template" template="https://sandbox.familysearch.org/platform/tree/persons/{pid}/child-relationships{?persons,access_token}" title="Relationships to Children" type="application/json,application/x-fs-v1+json,application/x-fs-v1+xml,application/x-gedcomx-v1+json,application/x-gedcomx-v1+xml,application/xml,text/html" accept="*/*" allow="GET"></atom:link>
</atom:feed>
		`)
	})

	want := map[string]string{
		"agent-template":                  "https://sandbox.familysearch.org/platform/users/agents/{uid}",
		"child-and-parents-relationships": fstesting.Server.URL + "/platform/tree/child-and-parents-relationships",
		"child-relationships-template":    "https://sandbox.familysearch.org/platform/tree/persons/{pid}/child-relationships",
	}
	client := &http.Client{}
	templates, err := readTemplates(fstesting.Server.URL, client)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(templates, want) {
		t.Errorf("readTemplates returned %+v, want %+v", templates, want)
	}
}
