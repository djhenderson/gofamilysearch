package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"net/http"
)

func TestGetOrdinanceAccess(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetOrdinanceAccess", t, func() {
		testRespond(t, "GET", "/platform/ordinances/ordinances", nil, "")
		access, err := testClient().GetOrdinanceAccess()
		So(err, ShouldBeNil)
		So(access, ShouldBeTrue)
	})
}

func TestGetOrdinanceAccessForbidden(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetOrdinanceAccessForbidden", t, func() {
		testMux.HandleFunc("/platform/ordinances/ordinances", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.WriteHeader(403)
		})
		access, err := testClient().GetOrdinanceAccess()
		So(err, ShouldBeNil)
		So(access, ShouldBeFalse)
	})
}
