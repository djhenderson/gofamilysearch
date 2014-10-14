package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReadDiscoveryResource(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("ReadDiscoveryResource", t, func() {
		testRespond(t, "GET", "/.well-known/app-meta", nil, "")
		templates, err := testClient().readDiscoveryResource(testServer.URL)
		So(err, ShouldBeNil)
		want := map[string]string{
			"person-template":                 "https://sandbox.familysearch.org/platform/tree/persons/{pid}",
			"person-with-relationships-query": "https://sandbox.familysearch.org/platform/tree/persons-with-relationships",
			"persons":                         "https://sandbox.familysearch.org/platform/tree/persons",
		}
		for k, v := range want {
			So(templates[k], ShouldEqual, v)
		}
	})
}
