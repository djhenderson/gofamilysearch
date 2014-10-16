package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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
