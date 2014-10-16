package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPersonChanges(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonChanges", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/P12-345/changes", nil, "")
		changes, err := testClient().GetPersonChanges("P12-345", 0, "")
		So(err, ShouldBeNil)
		wantChange := &Change{
			ID:      "1386263928318",
			Updated: 1386263928318,
			Title:   "Person Created",
			Links: map[string]*FSHref{
				"restore": &FSHref{Href: "https://familysearch.org/platform/tree/changes/153014/restore"},
				"agent":   &FSHref{Href: "https://familysearch.org/platform/users/agents/UKMGTY"},
			},
			Contributors: []*ChangeContributor{
				&ChangeContributor{Name: "Mr. Contributor"},
			},
			ChangeInfo: []*ChangeInfo{
				&ChangeInfo{
					Operation:  "http://familysearch.org/v1/Create",
					ObjectType: "http://gedcomx.org/Person",
					Reason:     "because it was necessary",
				},
			},
		}
		So(len(changes), ShouldEqual, 3)
		So(changes[0], ShouldResemble, wantChange)
	})
}
