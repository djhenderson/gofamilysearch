package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetHistoricalRecord(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetHistoricalRecord", t, func() {
		testRespond(t, "GET", "/ark:/61903/1:1:MPXD-MZC", nil, "historical_record.json")
		record, err := testClient().GetHistoricalRecord(testServer.URL + "/ark:/61903/1:1:MPXD-MZC")
		So(err, ShouldBeNil)

		wantSourceDescription := &HistoricalRecordSourceDescription{
			ID:           "src_r_1153479342",
			About:        "https://familysearch.org/ark:/61903/1:2:9JX9-6LD",
			ResourceType: "http://gedcomx.org/Record",
			Created:      1383765808862,
			Modified:     1413073490036,
			Citations: []*FSValue{&FSValue{
				Value: "\"United States Census, 1910,\" index and images, <i>FamilySearch</i> (https://familysearch.org/ark:/61903/1:2:9JX9-6LD : accessed 15 Oct 2014), Household of John W Spriggs, Prairie Belle, Beadle, South Dakota, United States; citing enumeration district (ED) 12, sheet 4A, family 63, NARA microfilm publication T624, FHL microfilm 1375489.",
			}},
			Titles: []*FSValue{&FSValue{
				Value: "Household of John W Spriggs, \"United States Census, 1910\"",
				Lang:  "en-US",
			}},
			Identifiers: map[string][]string{
				"http://gedcomx.org/Persistent": []string{"https://familysearch.org/ark:/61903/1:2:9JX9-6LD"},
			},
			ComponentOf: FSDescription{Description: "#src_2"},
			Sources: []*FSDescription{
				&FSDescription{Description: "#src_s1"},
				&FSDescription{Description: "#src_s2"},
			},
			Coverage: []*HistoricalRecordCoverage{
				&HistoricalRecordCoverage{
					RecordType: "http://gedcomx.org/Census",
					Temporal:   HistoricalRecordTemporalCoverage{Formal: "+1910", Original: "1910"},
					Spatial:    HistoricalRecordSpatialCoverage{Original: "Prairie Belle, Beadle, South Dakota, United States"},
				},
			},
		}

		wantPerson := &HistoricalRecordPerson{
			ID:        "p_570564307",
			Extracted: true,
			Principal: true,
			Gender: HistoricalRecordGender{
				Type: "http://gedcomx.org/Male",
				Fields: []*HistoricalRecordField{
					&HistoricalRecordField{
						Type: "http://gedcomx.org/Gender",
						Values: []*HistoricalRecordFieldValue{
							&HistoricalRecordFieldValue{
								Resource: "http://gedcomx.org/Unknown",
								Text:     "Male",
								LabelID:  "PR_SEX_CODE_ORIG",
								Type:     "http://gedcomx.org/Original",
							},
							&HistoricalRecordFieldValue{
								Resource: "http://gedcomx.org/Male",
								Text:     "Male",
								LabelID:  "PR_SEX_CODE",
								Type:     "http://gedcomx.org/Interpreted",
							},
						},
					},
				},
			},
			Names: []*HistoricalRecordName{
				&HistoricalRecordName{
					Type: "http://gedcomx.org/BirthName",
					NameForms: []*HistoricalRecordNameForm{
						&HistoricalRecordNameForm{
							FullText: "John W Spriggs",
							Parts: []*HistoricalRecordNamePart{
								&HistoricalRecordNamePart{
									NamePart: NamePart{
										Value: "John W",
										Type:  "http://gedcomx.org/Given",
									},
									Fields: []*HistoricalRecordField{
										&HistoricalRecordField{
											Type: "http://gedcomx.org/Given",
											Values: []*HistoricalRecordFieldValue{
												&HistoricalRecordFieldValue{
													Text:    "John W",
													LabelID: "PR_NAME_GN_ORIG",
													Type:    "http://gedcomx.org/Original",
												},
												&HistoricalRecordFieldValue{
													Text:    "John W",
													LabelID: "PR_NAME_GN",
													Type:    "http://gedcomx.org/Interpreted",
												},
											},
										},
									},
								},
								&HistoricalRecordNamePart{
									NamePart: NamePart{
										Value: "Spriggs",
										Type:  "http://gedcomx.org/Surname",
									},
									Fields: []*HistoricalRecordField{
										&HistoricalRecordField{
											Type: "http://gedcomx.org/Surname",
											Values: []*HistoricalRecordFieldValue{
												&HistoricalRecordFieldValue{
													Text:    "Spriggs",
													LabelID: "PR_NAME_SURN_ORIG",
													Type:    "http://gedcomx.org/Original",
												},
												&HistoricalRecordFieldValue{
													Text:    "Spriggs",
													LabelID: "PR_NAME_SURN",
													Type:    "http://gedcomx.org/Interpreted",
												},
											},
										},
									},
								},
							},
							Fields: []*HistoricalRecordField{
								&HistoricalRecordField{
									Type: "http://gedcomx.org/Name",
									Values: []*HistoricalRecordFieldValue{
										&HistoricalRecordFieldValue{
											Text:    "John W/Spriggs",
											LabelID: "PR_NAME_ORIG",
											Type:    "http://gedcomx.org/Original",
										},
										&HistoricalRecordFieldValue{
											Text:    "John W Spriggs",
											LabelID: "PR_NAME",
											Type:    "http://gedcomx.org/Interpreted",
										},
									},
								},
							},
						},
					},
				},
			},
			Facts: []*HistoricalRecordFact{
				&HistoricalRecordFact{
					Type:  "http://gedcomx.org/MaritalStatus",
					Value: "Married",
					Fields: []*HistoricalRecordField{
						&HistoricalRecordField{
							Type: "http://gedcomx.org/MaritalStatus",
							Values: []*HistoricalRecordFieldValue{
								&HistoricalRecordFieldValue{
									Text:    "Married",
									LabelID: "PR_MARITAL_STATUS_ORIG",
									Type:    "http://gedcomx.org/Original",
								},
								&HistoricalRecordFieldValue{
									Text:    "Married",
									LabelID: "PR_MARITAL_STATUS",
									Type:    "http://gedcomx.org/Interpreted",
								},
							},
						},
					},
				},
				&HistoricalRecordFact{
					Type:    "http://gedcomx.org/Census",
					Primary: true,
					Date: HistoricalRecordDate{
						Date: Date{
							Original: "1910",
						},
						Fields: []*HistoricalRecordField{
							&HistoricalRecordField{
								Type: "http://gedcomx.org/Date",
								Values: []*HistoricalRecordFieldValue{
									&HistoricalRecordFieldValue{
										Text:    "1910",
										LabelID: "EVENT_DATE",
										Type:    "http://gedcomx.org/Interpreted",
									},
								},
							},
							&HistoricalRecordField{
								Type: "http://gedcomx.org/Year",
								Values: []*HistoricalRecordFieldValue{
									&HistoricalRecordFieldValue{
										Text:    "1910",
										LabelID: "EVENT_YEAR",
										Type:    "http://gedcomx.org/Interpreted",
									},
								},
							},
						},
					},
					Place: HistoricalRecordPlace{
						Place: Place{
							Original: "Prairie Belle, Beadle, South Dakota, United States",
						},
						Fields: []*HistoricalRecordField{
							&HistoricalRecordField{
								Type: "http://gedcomx.org/Place",
								Values: []*HistoricalRecordFieldValue{
									&HistoricalRecordFieldValue{
										Text:    "Prairie Belle, Beadle, South Dakota, United States",
										LabelID: "EVENT_PLACE_ORIG",
										Type:    "http://gedcomx.org/Original",
									},
									&HistoricalRecordFieldValue{
										Text:    "Prairie Belle, Beadle, South Dakota, United States",
										LabelID: "EVENT_PLACE",
										Type:    "http://gedcomx.org/Interpreted",
									},
								},
							},
						},
					},
				},
				&HistoricalRecordFact{
					Type: "http://gedcomx.org/Birth",
					Date: HistoricalRecordDate{
						Date: Date{
							Original: "1872",
						},
						Fields: []*HistoricalRecordField{
							&HistoricalRecordField{
								Type: "http://gedcomx.org/Year",
								Values: []*HistoricalRecordFieldValue{
									&HistoricalRecordFieldValue{
										Text:    "1872",
										LabelID: "PR_BIRTH_YEAR_ESTIMATED",
										Type:    "http://gedcomx.org/Interpreted",
									},
								},
							},
						},
					},
					Place: HistoricalRecordPlace{
						Place: Place{
							Original: "Iowa",
						},
						Fields: []*HistoricalRecordField{
							&HistoricalRecordField{
								Type: "http://gedcomx.org/Place",
								Values: []*HistoricalRecordFieldValue{
									&HistoricalRecordFieldValue{
										Text:    "Iowa",
										LabelID: "PR_BIRTH_PLACE",
										Type:    "http://gedcomx.org/Interpreted",
									},
								},
							},
						},
					},
				},
			},
			Links: map[string]*FSHref{
				"persona": &FSHref{Href: "https://familysearch.org/platform/records/personas/MPXD-MZ4"},
				"record":  &FSHref{Href: "https://familysearch.org/platform/records/records/9JX9-6LD"},
			},
			Identifiers: map[string][]string{
				"$": []string{"https://familysearch.org/platform/externalId/easy/1046930057"},
				"http://gedcomx.org/Persistent": []string{"https://familysearch.org/ark:/61903/1:1:MPXD-MZ4"},
			},
			Fields: []*HistoricalRecordField{
				&HistoricalRecordField{
					Type: "http://gedcomx.org/Age",
					Values: []*HistoricalRecordFieldValue{
						&HistoricalRecordFieldValue{
							Text:    "38y",
							LabelID: "PR_AGE_ORIG",
							Type:    "http://gedcomx.org/Original",
						},
						&HistoricalRecordFieldValue{
							Text:    "38",
							LabelID: "PR_AGE",
							Type:    "http://gedcomx.org/Interpreted",
						},
					},
				},
				&HistoricalRecordField{
					Type: "http://gedcomx.org/Race",
					Values: []*HistoricalRecordFieldValue{
						&HistoricalRecordFieldValue{
							Text:    "White",
							LabelID: "PR_RACE_OR_COLOR_ORIG",
							Type:    "http://gedcomx.org/Original",
						},
						&HistoricalRecordFieldValue{
							Text:    "White",
							LabelID: "PR_RACE_OR_COLOR",
							Type:    "http://gedcomx.org/Interpreted",
						},
					},
				},
			},
		}

		wantRelationship := &HistoricalRecordRelationship{
			ID:      "MM9.1.6/SVS9-TW3",
			Type:    "http://gedcomx.org/Couple",
			Person1: ResourceRef{Resource: "#p_570564307"},
			Person2: ResourceRef{Resource: "#p_570564308"},
		}

		So(record.Description, ShouldEqual, "#src_p_570564309")
		So(len(record.SourceDescriptions), ShouldEqual, 11)
		So(record.SourceDescriptions[0], ShouldResemble, wantSourceDescription)
		So(len(record.Persons), ShouldEqual, 6)
		So(record.Persons[0], ShouldResemble, wantPerson)
		So(len(record.Relationships), ShouldEqual, 7)
		So(record.Relationships[0], ShouldResemble, wantRelationship)
		So(record.ID, ShouldEqual, "r_1153479342")
	})
}

func TestGetRecordCollection(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetRecordCollection", t, func() {
		testRespond(t, "GET", "/platform/records/collections/1727033", nil, "historical_collection.json")
		collection, err := testClient().GetRecordCollection(testServer.URL + "/platform/records/collections/1727033")
		So(err, ShouldBeNil)

		wantSourceDescription := &RecordCollectionSourceDescription{
			ID:    "1727033",
			About: "https://familysearch.org/platform/records/collections/1727033",
			Citations: []*FSValue{&FSValue{
				Value: "\"United States Census, 1910.\" Index and images. <i>FamilySearch</i>. http://FamilySearch.org : accessed 2014. Citing NARA microfilm publication T624. Washington, D.C.: National Archives and Records Administration, n.d.",
			}},
			ComponentOf: FSDescription{Description: "#1727033_c"},
			Titles: []*FSValue{&FSValue{
				Value: "United States Census, 1910",
			}},
			ResourceType: "http://gedcomx.org/Collection",
			Rights: []string{
				"http://familysearch.org/records/permissionGroup/FamilySearch",
				"http://familysearch.org/records/permissionGroup/Partner2",
			},
			Descriptions: []*FSValue{&FSValue{
				Value: "Index to the 1910 population census schedules. Indexing is currently in progress and will include the entire census comprising 48 states, two territories (Arizona and New Mexico), Puerto Rico, and Military and Naval (in Philippines, Hospitals, Ships, and Stations). The index is being created by FamilySearch and Ancestry.com.",
				Lang:  "en_US",
			}},
			Identifiers: map[string][]string{
				"http://gedcomx.org/Primary": []string{"https://familysearch.org/platform/records/collections/1727033"},
			},
			Coverage: []*HistoricalRecordCoverage{
				&HistoricalRecordCoverage{
					RecordType: "http://gedcomx.org/Census",
					Temporal:   HistoricalRecordTemporalCoverage{Formal: "+1910", Original: "1910"},
					Spatial:    HistoricalRecordSpatialCoverage{Original: "United States of America"},
				},
			},
		}

		wantCollection := &RecordCollectionCollection{
			Lang: "en-US",
			Content: []*RecordCollectionContent{
				&RecordCollectionContent{
					ResourceType: "http://gedcomx.org/Record",
					Count:        93682083,
					Completeness: 0.98,
				},
				&RecordCollectionContent{
					ResourceType: "http://gedcomx.org/Person",
					Count:        93682083,
					Completeness: 0.98,
				},
				&RecordCollectionContent{
					ResourceType: "http://gedcomx.org/DigitalArtifact",
					Count:        2068253,
					Completeness: 0,
				},
			},
			Title: "United States Census, 1910",
			Size:  95750336,
		}

		wantLabel := &FSValue{
			Lang:  "en_US",
			Value: "Name",
		}

		So(collection.Description, ShouldEqual, "#1727033")
		So(len(collection.SourceDescriptions), ShouldEqual, 2)
		So(collection.SourceDescriptions[0], ShouldResemble, wantSourceDescription)
		So(len(collection.Collections), ShouldEqual, 1)
		So(collection.Collections[0], ShouldResemble, wantCollection)
		So(len(collection.RecordDescriptors), ShouldEqual, 3)
		So(collection.RecordDescriptors[0].ID, ShouldEqual, "rd_MM9.1.9/MMMM-M5N")
		So(len(collection.RecordDescriptors[0].Fields), ShouldBeGreaterThan, 0)
		So(len(collection.RecordDescriptors[0].Fields[0].Values), ShouldBeGreaterThan, 0)
		So(collection.RecordDescriptors[0].Fields[0].Values[0].LabelID, ShouldEqual, "PR_NAME")
		So(len(collection.RecordDescriptors[0].Fields[0].Values[0].Labels), ShouldBeGreaterThan, 0)
		So(collection.RecordDescriptors[0].Fields[0].Values[0].Labels[0], ShouldResemble, wantLabel)
	})
}
