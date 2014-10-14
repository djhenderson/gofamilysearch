package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPersonNotes(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonNotes", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/P12-3456/notes", nil, "")
		notes, err := testClient().GetPersonNotes("P12-3456")
		So(err, ShouldBeNil)
		want := []*Note{
			&Note{ID: "1804317705",
				Subject: "note 0",
				Text: "Sample note text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
			&Note{ID: "1805241226",
				Subject: "note 1",
				Text: "text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
		}
		So(notes, ShouldResemble, want)
	})
}

func TestGetCoupleNotes(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetCoupleNotes", t, func() {
		testRespond(t, "GET", "/platform/tree/couple-relationships/12345/notes", nil, "")
		notes, err := testClient().GetCoupleNotes("12345")
		So(err, ShouldBeNil)
		want := []*Note{
			&Note{ID: "1804317705",
				Subject: "note 0",
				Text: "text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
			&Note{ID: "1805241226",
				Subject: "note 1",
				Text: "text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
		}
		So(notes, ShouldResemble, want)
	})
}

func TestGetChildAndParentsNotes(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetChildAndParentsNotes", t, func() {
		testRespond(t, "GET", "/platform/tree/child-and-parents-relationships/PPPX-PP0/notes", nil, "")
		notes, err := testClient().GetChildAndParentsNotes("PPPX-PP0")
		So(err, ShouldBeNil)
		want := []*Note{
			&Note{ID: "1804317705",
				Subject: "note 0",
				Text: "text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
			&Note{ID: "1805241226",
				Subject: "note 1",
				Text: "text",
				Attribution: Attribution{
					Contributor: ResourceRef{
						Resource: "https://sandbox.familysearch.org/platform/users/agents/MMD8-3NT",
						ResourceID: "MMD8-3NT"},
					Modified: 1403312322000}},
		}
		So(notes, ShouldResemble, want)
	})
}

