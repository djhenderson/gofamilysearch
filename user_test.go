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
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<fs:familysearch xmlns="http://gedcomx.org/v1/" xmlns:fs="http://familysearch.org/v1/" xmlns:atom="http://www.w3.org/2005/Atom">
    <fs:user id="cis.MMM.RX9">
        <link rel="self" href="https://familysearch.org/platform/users/current"/>
        <link rel="person" href="https://familysearch.org/platform/tree/persons/JNM-VRQM"/>
        <fs:contactName>Pete Townsend</fs:contactName>
        <fs:email>peter@acme.org</fs:email>
        <fs:personId>JNM-VRQM</fs:personId>
        <fs:treeUserId>PXRQ-FMXT</fs:treeUserId>
    </fs:user>
</fs:familysearch>
		`)
	})

	want := &User{
		Id : "cis.MMM.RX9",
		ContactName : "Pete Townsend",
		Email : "peter@acme.org",
		PersonId : "JNM-VRQM",
		TreeUserId : "PXRQ-FMXT",
	}

	c := &Client{
		AccessToken: "accessToken",
		HttpClient: &http.Client{},
		Context: &Context{
			templates:   map[string]string{
				"current-user" : fstesting.Server.URL + "/current-user",
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
