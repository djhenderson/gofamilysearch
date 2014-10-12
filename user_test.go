package gofamilysearch

import (
	"reflect"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	testSetup()
	defer testTeardown()

	testRespond(t, "GET", "/platform/users/current", nil, "platform_users_current.json")

	user, err := testClient().GetCurrentUser()
	if err != nil {
		t.Errorf("GetCurrentUser error %s", err.Error())
	}

	want := &User{
		Id:          "cis.MMM.RX9",
		ContactName: "Pete Townsend",
		Email:       "peter@acme.org",
		TreeUserId:  "PXRQ-FMXT",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("GetCurrentUser returned %+v, want %+v", user)
	}
}
