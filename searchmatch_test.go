package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPersonMatches(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonMatches", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/KF5C-9TX/matches",
			map[string]string{"confidence": "2", "count": "100", "collection": "https://familysearch.org/platform/collections/records"},
			"platform_tree_persons_KF5C-9TX_matches_collection_https___familysearch_org_platform_collections_records_count_100_confidence_2.json")
		matches, err := testClient().GetPersonMatches("KF5C-9TX", "https://familysearch.org/platform/collections/records", 2, 100)
		So(err, ShouldBeNil)
		want := []*SearchMatchResult{
			&SearchMatchResult{
				ID:         "https://familysearch.org/ark:/61903/1:1:MPXD-MZC",
				Score:      3.6129,
				Confidence: 5,
				Links: map[string]*FSHref{
					"source-linker": &FSHref{Href: "https://beta.familysearch.org/platform/redirect?person=KF5C-9TX&context=sourceLinker&hintId=https://familysearch.org/ark:/61903/1:1:MPXD-MZC"},
					"person":        &FSHref{Href: "https://beta.familysearch.org/platform/tree/persons/KF5C-9TX"},
					"record-hints":  &FSHref{Href: "https://beta.familysearch.org/platform/redirect?person=KF5C-9TX&context=recordHints"},
				},
				Published: 1405999060485,
				Title:     "United States Census, 1910",
				MatchInfo: []*MatchInfo{
					&MatchInfo{
						Collection: "https://familysearch.org/platform/collections/records",
						Status:     "http://familysearch.org/v1/Pending",
					},
				},
			},
		}
		So(matches, ShouldResemble, want)
	})
}
