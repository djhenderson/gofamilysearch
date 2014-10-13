package gofamilysearch

import (
	. "github.com/smartystreets/goconvey/convey"
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
