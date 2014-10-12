package gofamilysearch

import (
	"testing"
	"log"
)

func TestGetPersonWithRelationships(t *testing.T) {
	testSetup()
	defer testTeardown()

	testRespond(t, "GET", "/platform/tree/persons-with-relationships",
		map[string]string{"person": "KW7S-VQJ", "persons": "true"},
		"platform_tree_persons-with-relationships_person_KW7S-VQJ_persons_true.json")

	pwr, err := testClient().GetPersonWithRelationships("KW7S-VQJ")
	if err != nil {
		t.Errorf("GetPersonWithRelationships error %s", err.Error())
	}
	log.Printf("PWR=%#v\n", pwr)
}
