package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestGetPersonWithRelationships(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonWithRelationships", t, func() {
		id := "KW7S-VQJ"
		testRespond(t, "GET", "/platform/tree/persons-with-relationships",
			map[string]string{"person": "KW7S-VQJ", "persons": "true"},
			"platform_tree_persons-with-relationships_person_"+id+"_persons_true.json")
		pwr, err := testClient().GetPersonWithRelationships(id)
		So(err, ShouldBeNil)
		So(pwr.GetPerson("KW7S-VQJ").ID, ShouldEqual, id)
		So(len(pwr.GetParentRelationships(id)), ShouldBeGreaterThan, 0)
		So(pwr.GetPerson(pwr.GetParentRelationships(id)[0].Father.ResourceID).Display.Name,
			ShouldEqual, "Jens Christian Jensen")
		So(pwr.GetPerson(pwr.GetParentRelationships(id)[0].Mother.ResourceID).Display.Name,
			ShouldEqual, "Ane Christensdr")
		So(len(pwr.GetChildRelationships(id)), ShouldBeGreaterThan, 0)
		So(pwr.GetPerson(pwr.GetChildRelationships(id)[0].Child.ResourceID).Display.Name,
			ShouldEqual, "Christian Ludvic Jensen")
		So(len(pwr.GetSpouseRelationships()), ShouldBeGreaterThan, 0)
		So(pwr.GetSpouseRelationships()[0].Person2.ResourceID, ShouldEqual, "KW7S-JB7")
	})
}

func TestGetPersonPortraitURL(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetPersonPortraitURL", t, func() {
		testMux.HandleFunc("/platform/tree/persons/123/portrait", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.Header().Set("Location", "http://familysearch.org/pic.jpg")
			w.WriteHeader(307)
		})
		url, err := testClient().GetPersonPortraitURL("123")
		So(err, ShouldBeNil)
		So(url, ShouldEqual, "http://familysearch.org/pic.jpg")
	})
}

func TestGetPersonPortraitURLNotFound(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetOrdinanceAccessForbidden", t, func() {
		testMux.HandleFunc("/platform/tree/persons/123/portrait", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.WriteHeader(204)
		})
		url, err := testClient().GetPersonPortraitURL("123")
		So(err, ShouldBeNil)
		So(url, ShouldEqual, "")
	})
}

func TestGetPreferredParentsID(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetGetPreferredParentsID", t, func() {
		testMux.HandleFunc("/platform/tree/users/PXRQ-FMXT/preferred-parent-relationships/PPPJ-MYY", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.Header().Set("Location", "https://familysearch.org/platform/tree/child-and-parents-relationships/23456")
			w.WriteHeader(303)
		})
		id, err := testClient().GetPreferredParentsID("PXRQ-FMXT", "PPPJ-MYY")
		So(err, ShouldBeNil)
		So(id, ShouldEqual, "23456")
	})
}

func TestGetPreferredParentsIDNone(t *testing.T) {
	testSetup()
	defer testTeardown()

	Convey("GetGetPreferredParentsIDNone", t, func() {
		testMux.HandleFunc("/platform/tree/users/PXRQ-FMXT/preferred-parent-relationships/PPPJ-MYY", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.WriteHeader(204)
		})
		id, err := testClient().GetPreferredParentsID("PXRQ-FMXT", "PPPJ-MYY")
		So(err, ShouldBeNil)
		So(id, ShouldEqual, "")
	})
}
