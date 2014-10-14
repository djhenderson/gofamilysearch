package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetDiscussion(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetDiscussion", t, func() {
		testRespond(t, "GET", "/platform/discussions/discussions/dis-MMMM-MMM", nil, "")
		discussion, err := testClient().GetDiscussion("dis-MMMM-MMM")
		So(err, ShouldBeNil)
		want := &Discussion{
			ID:               "dis-MMMM-MMM",
			Title:            "1900 US Census, Ethel Hollivet",
			Details:          "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  ",
			NumberOfComments: 2,
			Contributor: ResourceRef{
				Resource:   "https://sandbox.familysearch.org/platform/users/agents/12345",
				ResourceID: "12345"},
		}
		So(discussion, ShouldResemble, want)
	})
}

func TestGetPersonDiscussionRefs(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonDiscussionRefs", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/12345/discussion-references", nil, "")
		discussionRefs, err := testClient().GetPersonDiscussionRefs("12345")
		So(err, ShouldBeNil)
		want := []*DiscussionRef{
			&DiscussionRef{
				Resource: "https://familysearch.org/platform/discussions/discussions/dis-MMMM-MMM",
			},
		}
		So(discussionRefs, ShouldResemble, want)
	})
}

func TestGetDiscussionComments(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetDiscussionComments", t, func() {
		testRespond(t, "GET", "/platform/discussions/discussions/dis-MMMM-MMM/comments", nil, "")
		comments, err := testClient().GetDiscussionComments("dis-MMMM-MMM")
		So(err, ShouldBeNil)
		want := []*Comment{
			&Comment{
				ID:   "CMMM-MMM",
				Text: "Just a comment.",
			},
		}
		So(comments, ShouldResemble, want)
	})
}
