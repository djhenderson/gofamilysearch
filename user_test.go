package gofamilysearch

import (
	"fmt"
	"github.com/rootsdev/gofamilysearch/fstesting"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	fstesting.Setup()
	defer fstesting.Teardown()

	fstesting.Mux.HandleFunc("/current-user", func(w http.ResponseWriter, r *http.Request) {
		fstesting.TestMethod(t, r, "GET")
		fmt.Fprintf(w, `
{
  "users" : [ {
    "id" : "cis.MMM.RX9",
    "links" : {
      "person" : {
        "href" : "https://familysearch.org/platform/tree/persons/JNM-VRQM"
      },
      "self" : {
        "href" : "https://familysearch.org/platform/users/current"
      }
    },
    "contactName" : "Pete Townsend",
    "fullName" : "Pete Townsend",
    "email" : "peter@acme.org",
    "personId" : "JNM-VRQM",
    "treeUserId" : "PXRQ-FMXT"
  } ]
}
		`)
	})

	want := &User{
		Id:          "cis.MMM.RX9",
		ContactName: "Pete Townsend",
		Email:       "peter@acme.org",
		PersonId:    "JNM-VRQM",
		TreeUserId:  "PXRQ-FMXT",
	}

	c := &Client{
		AccessToken: "accessToken",
		HttpClient:  &http.Client{},
		Context: 	 &Context{
			templates: map[string]string{
				"current-user": fstesting.Server.URL + "/current-user",
			},
		},
	}
	user, err := c.GetCurrentUser()
	if err != nil {
		t.Errorf("GetCurrentUser error %s", err.Error())
	}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("GetCurrentUser returned %+v, want %+v", user, want)
	}
}
