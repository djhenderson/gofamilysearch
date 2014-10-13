package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReadDiscoveryResource(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("ReadDiscoveryResource", t, func() {
		testRespond(t, "GET", "/platform/collections/tree", nil, "platform_collections_tree.json")
		templates, err := testClient().readDiscoveryResource(testServer.URL)
		So(err, ShouldBeNil)
		want := map[string]string{
			"person":                    "https://sandbox.familysearch.org/platform/tree/persons/{pid}",
			"person-with-relationships": "https://sandbox.familysearch.org/platform/tree/persons-with-relationships",
			"persons":                   "https://sandbox.familysearch.org/platform/tree/persons",
		}
		for k, v := range want {
			So(templates[k], ShouldEqual, v)
		}
	})
}
