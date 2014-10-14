package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPersonSources(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonSources", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/PPPP-PPP/sources", nil, "")
		sources, err := testClient().GetPersonSources("PPPP-PPP")
		So(err, ShouldBeNil)
		want := &Sources{
			SourceRefs: []*SourceRef{
				&SourceRef{
					ID: "abcde",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU"},
						Modified:      123456789,
						ChangeMessage: "Family is at the same address found in other sources associated with this family.  Names are a good match.  Estimated births are reasonable."},
					Description: "#SSSS-SS1",
					Tags: []*Tag{
						&Tag{
							Resource: "http://gedcomx.org/Name",
						},
						&Tag{
							Resource: "http://gedcomx.org/Gender",
						},
						&Tag{
							Resource: "http://gedcomx.org/Birth",
						},
					},
				},
				&SourceRef{
					ID: "fghij",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU",
						},
						Modified:      987654321,
						ChangeMessage: "Dates and location match with other sources.",
					},
					Description: "#SSSS-SS2",
				},
			},
			SourceDescriptions: []*SourceDescription{
				&SourceDescription{
					ID: "SSSS-SS1",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022617635,
						ChangeMessage: "This is the change message",
					},
				},
				&SourceDescription{
					ID: "SSSS-SS2",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022617635,
						ChangeMessage: "This is the change message",
					},
				},
			},
		}
		// So(len(sources.SourceRefs), ShouldEqual, len(want.SourceRefs))
		// for i, source := range sources.SourceRefs {
		//     So(source, ShouldResemble, want.SourceRefs[i])
		// }
		// So(len(sources.SourceDescriptions), ShouldEqual, len(want.SourceDescriptions))
		// for i, source := range sources.SourceDescriptions {
		//     So(source, ShouldResemble, want.SourceDescriptions[i])
		// }
		So(sources, ShouldResemble, want)
	})
}

func TestGetCoupleSources(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetCoupleSources", t, func() {
		testRespond(t, "GET", "/platform/tree/couple-relationships/12345/sources", nil, "")
		sources, err := testClient().GetCoupleSources("12345")
		So(err, ShouldBeNil)
		want := &Sources{
			SourceRefs: []*SourceRef{
				&SourceRef{
					ID: "abcde",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU"},
						Modified:      123456789,
						ChangeMessage: "Family is at the same address found in other sources associated with this family.  Names are a good match.  Estimated births are reasonable."},
					Description: "#SSSS-SS1",
					Tags: []*Tag{
						&Tag{
							Resource: "http://gedcomx.org/Name",
						},
						&Tag{
							Resource: "http://gedcomx.org/Gender",
						},
						&Tag{
							Resource: "http://gedcomx.org/Birth",
						},
					},
				},
				&SourceRef{
					ID: "fghij",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU",
						},
						Modified:      987654321,
						ChangeMessage: "Dates and location match with other sources.",
					},
					Description: "#SSSS-SS2",
				},
			},
			SourceDescriptions: []*SourceDescription{
				&SourceDescription{
					ID: "SSSS-SS1",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022713544,
						ChangeMessage: "This is the change message",
					},
				},
				&SourceDescription{
					ID: "SSSS-SS2",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022713544,
						ChangeMessage: "This is the change message",
					},
				},
			},
		}
		So(sources, ShouldResemble, want)
	})
}

func TestGetChildAndParentsSources(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetChildAndParentsSources", t, func() {
		testRespond(t, "GET", "/platform/tree/child-and-parents-relationships/PPPX-PP0/sources", nil, "")
		sources, err := testClient().GetChildAndParentsSources("PPPX-PP0")
		So(err, ShouldBeNil)
		want := &Sources{
			SourceRefs: []*SourceRef{
				&SourceRef{
					ID: "fghij",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU",
						},
						Modified:      987654321,
						ChangeMessage: "Dates and location match with other sources.",
					},
					Description: "#SSSS-SS2",
				},
				&SourceRef{
					ID: "abcde",
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/UUUU-UUU",
							ResourceID: "UUUU-UUU"},
						Modified:      123456789,
						ChangeMessage: "Family is at the same address found in other sources associated with this family.  Names are a good match.  Estimated births are reasonable."},
					Description: "#SSSS-SS1",
					Tags: []*Tag{
						&Tag{
							Resource: "http://gedcomx.org/Name",
						},
						&Tag{
							Resource: "http://gedcomx.org/Gender",
						},
						&Tag{
							Resource: "http://gedcomx.org/Birth",
						},
					},
				},
			},
			SourceDescriptions: []*SourceDescription{
				&SourceDescription{
					ID: "SSSS-SS2",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022662152,
						ChangeMessage: "This is the change message",
					},
				},
				&SourceDescription{
					ID: "SSSS-SS1",
					Citations: []*SourceCitation{
						&SourceCitation{
							Value: "\"United States Census, 1900.\" database and digital images, FamilySearch (https://familysearch.org/: accessed 17 Mar 2012), Ethel Hollivet, 1900; citing United States Census Office, Washington, D.C., 1900 Population Census Schedules, Los Angeles, California, population schedule, Los Angeles Ward 6, Enumeration District 58, p. 20B, dwelling 470, family 501, FHL microfilm 1,240,090; citing NARA microfilm publication T623, roll 90.",
						},
					},
					About: "https://familysearch.org/pal:/MM9.1.1/M9PJ-2JJ",
					Titles: []*SourceTitle{
						&SourceTitle{
							Value: "1900 US Census, Ethel Hollivet",
						},
					},
					Notes: []*SourceNote{
						&SourceNote{
							Text: "Ethel Hollivet (line 75) with husband Albert Hollivet (line 74); also in the dwelling: step-father Joseph E Watkins (line 72), mother Lina Watkins (line 73), and grandmother -- Lina's mother -- Mary Sasnett (line 76).  Albert's mother and brother also appear on this page -- Emma Hollivet (line 68), and Eddie (line 69).",
						},
					},
					Attribution: Attribution{
						Contributor: ResourceRef{
							Resource:   "https://familysearch.org/platform/users/agents/123",
							ResourceID: "123"},
						Modified:      1396022662152,
						ChangeMessage: "This is the change message",
					},
				},
			},
		}
		So(sources, ShouldResemble, want)
	})
}
