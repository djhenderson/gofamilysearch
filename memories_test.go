package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPersonMemories(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonMemories", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/KWCR-JWS/memories", map[string]string{"start": "2", "count": "2"},
			"platform_tree_persons_KWCR-JWS_memories_count_2_start_2.json")
		memories, err := testClient().GetPersonMemories("KWCR-JWS", 2, 2)
		So(err, ShouldBeNil)
		want := []*Memory{
			&Memory{
				ID: "904106",
				Links: map[string]*FSHref{
					"memory": &FSHref{Href: "https://familysearch.org/platform/memories/memories/904106"},
				},
				About:        "https://familysearch.org/photos/images/904106",
				Titles:       []*FSValue{&FSValue{Value: "Missionary Portrait"}},
				Descriptions: []*FSValue{&FSValue{Value: "Alma Heaton while on a mission to Canada."}},
				ArtifactMetadata: []*ArtifactMetadata{
					&ArtifactMetadata{
						Filename:     "alma-mission.jpg",
						ArtifactType: "http://familysearch.org/v1/Image",
					},
				},
			},
			&Memory{
				ID: "1559654",
				Links: map[string]*FSHref{
					"memory": &FSHref{Href: "https://familysearch.org/platform/memories/memories/1559654"},
				},
				About:        "https://familysearch.org/photos/images/1559654",
				Titles:       []*FSValue{&FSValue{Value: "Record Player"}},
				Descriptions: []*FSValue{&FSValue{Value: "Alma Heaton using his favorite record player."}},
				ArtifactMetadata: []*ArtifactMetadata{
					&ArtifactMetadata{
						Filename:     "alma-record-player.jpg",
						ArtifactType: "http://familysearch.org/v1/Image",
					},
				},
			},
		}
		So(memories, ShouldResemble, want)
	})
}

func TestGetPersonMemoryPersonaRefs(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonMemoryPersonaRefs", t, func() {
		testRespond(t, "GET", "/platform/tree/persons/PPPP-PPP/memory-references", nil, "")
		memoryPersonaRefs, err := testClient().GetPersonMemoryPersonaRefs("PPPP-PPP")
		So(err, ShouldBeNil)
		want := []*MemoryPersonaRef{
			&MemoryPersonaRef{
				Resource:   "https://familysearch.org/platform/memories/memories/ARXX-MMM/personas/1083",
				ResourceID: "1083",
			},
			&MemoryPersonaRef{
				Resource:   "https://familysearch.org/platform/memories/memories/ARXX-MMM/personas/1204",
				ResourceID: "1204",
			},
		}
		So(memoryPersonaRefs, ShouldResemble, want)
	})
}

func TestGetMemoryPersonas(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetMemoryPersonas", t, func() {
		testRespond(t, "GET", "/platform/memories/memories/AR-1234/personas", nil, "")
		memoryPersonas, err := testClient().GetMemoryPersonas("AR-1234")
		So(err, ShouldBeNil)
		want := []*MemoryPersona{
			&MemoryPersona{
				ID: "123",
				Media: []*MediaRef{
					&MediaRef{
						ID: "T123",
						Qualifiers: []*MediaQualifier{
							&MediaQualifier{
								Name:  "http://gedcomx.org/RectangleRegion",
								Value: ".10,.10,1.00,1.00",
							},
						},
						Description: "https://familysearch.org/platform/memories/artifacts/132692/description",
					},
				},
				Names: []*MemoryPersonaName{
					&MemoryPersonaName{
						NameForms: []*MemoryPersonaNameForm{
							&MemoryPersonaNameForm{
								FullText: "Anastasia Aleksandrova",
							},
						},
					},
				},
				Display: MemoryPersonaDisplay{
					Name: "Anastasia Aleksandrova",
				},
			},
		}
		So(memoryPersonas, ShouldResemble, want)
	})
}

func TestGetMemoryComments(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetMemoryComments", t, func() {
		testRespond(t, "GET", "/platform/memories/memories/AR-1234/comments", nil, "")
		comments, err := testClient().GetMemoryComments("AR-1234")
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
