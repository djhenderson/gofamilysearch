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

	fstesting.Mux.HandleFunc("/.well-known/collections/tree", func(w http.ResponseWriter, r *http.Request) {
		fstesting.TestMethod(t, r, "GET")
		fmt.Fprintf(w, `
{
  "collections" : [ {
    "id" : "FSFT",
    "links" : {
      "agent-template" : {
        "template" : "https://sandbox.familysearch.org/platform/users/agents/{uid}{?access_token}"
      },
      "child-and-parents-relationships" : {
        "href" : "/platform/tree/child-and-parents-relationships"
      },
      "child-relationships-template" : {
        "template" : "https://sandbox.familysearch.org/platform/tree/persons/{pid}/child-relationships{?persons,access_token}"
      }
    }
  } ]
}
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
