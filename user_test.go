package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetCurrentUser", t, func() {
		testRespond(t, "GET", "/platform/users/current", nil, "platform_users_current.json")
		user, err := testClient().GetCurrentUser()
		So(err, ShouldBeNil)
		want := &User{
			ID:          "cis.MMM.RX9",
			ContactName: "Pete Townsend",
			Email:       "peter@acme.org",
			TreeUserID:  "PXRQ-FMXT",
		}
		So(user, ShouldResemble, want)
	})
}
