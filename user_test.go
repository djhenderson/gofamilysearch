package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetCurrentUser", t, func() {
		testRespond(t, "GET", "/platform/users/current", nil, "")
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

func TestGetAgent(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetAgent", t, func() {
		testRespond(t, "GET", "/platform/users/agents/12345", nil, "")
		agent, err := testClient().GetAgent("12345")
		So(err, ShouldBeNil)
		want := &Agent{
			ID: "12345",
			Names: []*AgentName{&AgentName{
				Value: "John Smith",
				Type:  "http://familysearch.org/v1//DisplayName",
			}},
			Accounts: []*AgentAccount{&AgentAccount{
				AccountName: "account",
			}},
			Emails: []*AgentEmail{&AgentEmail{
				Resource: "mailto:someone@somewhere.org",
			}},
		}
		So(agent, ShouldResemble, want)
	})
}
