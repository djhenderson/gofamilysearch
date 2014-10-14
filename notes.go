package gofamilysearch

import "fmt"

type personNotesResponse struct {
	Persons []*noteContainer `json:"persons"`
}

type coupleNotesResponse struct {
	Relationships []*noteContainer `json:"relationships"`
}

type childAndParentsNotesResponse struct {
	ChildAndParentsRelationships []*noteContainer `json:"childAndParentsRelationships"`
}

type noteContainer struct {
	Notes []*Note `json:"notes"`
}

// Note contains information about a note
type Note struct {
	ID          string      `json:"id"`
	Subject     string      `json:"subject"`
	Text        string      `json:"text"`
	Attribution Attribution `json:"attribution"`
}

// GetPersonNotes https://familysearch.org/developers/docs/api/tree/Person_Notes_resource
func (c *Client) GetPersonNotes(pid string) ([]*Note, error) {
	u, err := c.GetURL("person-notes-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	personNotesResponse := &personNotesResponse{}
	if err = c.Get(u, nil, nil, personNotesResponse); err != nil {
		return nil, err
	}
	if len(personNotesResponse.Persons) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return personNotesResponse.Persons[0].Notes, nil
}

// GetCoupleNotes https://familysearch.org/developers/docs/api/tree/Couple_Relationship_Notes_resource
func (c *Client) GetCoupleNotes(crid string) ([]*Note, error) {
	u, err := c.GetURL("couple-relationship-notes-template", map[string]string{"crid": crid})
	if err != nil {
		return nil, err
	}
	coupleNotesResponse := &coupleNotesResponse{}
	if err = c.Get(u, nil, nil, coupleNotesResponse); err != nil {
		return nil, err
	}
	if len(coupleNotesResponse.Relationships) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return coupleNotesResponse.Relationships[0].Notes, nil
}

// GetChildAndParentsNotes https://familysearch.org/developers/docs/api/tree/Child-and-Parents_Relationship_Notes_resource
func (c *Client) GetChildAndParentsNotes(caprid string) ([]*Note, error) {
	u, err := c.GetURL("child-and-parents-relationship-notes-template", map[string]string{"caprid": caprid})
	if err != nil {
		return nil, err
	}
	childAndParentsNotesResponse := &childAndParentsNotesResponse{}
	if err = c.Get(u, nil, nil, childAndParentsNotesResponse); err != nil {
		return nil, err
	}
	if len(childAndParentsNotesResponse.ChildAndParentsRelationships) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return childAndParentsNotesResponse.ChildAndParentsRelationships[0].Notes, nil
}
