package gofamilysearch

import (
	"fmt"
	"strings"
)

// PersonWithRelationships is the GetPersonWithRelationships response
type PersonWithRelationships struct {
	Persons                      []*Person                      `json:"persons"`
	ChildAndParentsRelationships []*ChildAndParentsRelationship `json:"childAndParentsRelationships"`
	Relationships                []*Relationship                `json:"relationships"`
}

// GetPerson get the person with id
func (pwr *PersonWithRelationships) GetPerson(id string) *Person {
	for _, person := range pwr.Persons {
		if person.ID == id {
			return person
		}
	}
	return nil
}

// GetParentRelationships gets the ChildAndParentsRelationships where id is a child
func (pwr *PersonWithRelationships) GetParentRelationships(id string) []*ChildAndParentsRelationship {
	rels := make([]*ChildAndParentsRelationship, 0, len(pwr.ChildAndParentsRelationships))
	for _, capr := range pwr.ChildAndParentsRelationships {
		if capr.Child.ResourceID == id {
			rels = append(rels, capr)
		}
	}
	return rels
}

// GetChildRelationships gets the ChildAndParentRelationships where id is a parent
func (pwr *PersonWithRelationships) GetChildRelationships(id string) []*ChildAndParentsRelationship {
	rels := make([]*ChildAndParentsRelationship, 0, len(pwr.ChildAndParentsRelationships))
	for _, capr := range pwr.ChildAndParentsRelationships {
		if capr.Father.ResourceID == id || capr.Mother.ResourceID == id {
			rels = append(rels, capr)
		}
	}
	return rels
}

// GetSpouseRelationships gets the Relationships that are couple relationships
func (pwr *PersonWithRelationships) GetSpouseRelationships() []*Relationship {
	rels := make([]*Relationship, 0, len(pwr.Relationships))
	for _, cr := range pwr.Relationships {
		if cr.Type == "http://gedcomx.org/Couple" {
			rels = append(rels, cr)
		}
	}
	return rels
}

// Person contains information about a person
type Person struct {
	ID          string              `json:"id"`
	Living      bool                `json:"living"`
	Display     PersonDisplay       `json:"display"`
	Identifiers map[string][]string `json:"identifiers"`
	Gender      Gender              `json:"gender"`
	Attribution Attribution         `json:"attribution"`
	Facts       []*Fact             `json:"facts"`
	Names       []*Name             `json:"names"`
	Links       map[string]*FSHref  `json:"links"`
}

// PersonDisplay contains various person attributes in ready-to-display format
type PersonDisplay struct {
	BirthDate  string `json:"birthDate"`
	BirthPlace string `json:"birthPlace"`
	DeathDate  string `json:"deathDate"`
	DeathPlace string `json:"deathPlace"`
	Gender     string `json:"gender"`
	LifeSpan   string `json:"lifespan"`
	Name       string `json:"name"`
}

// Gender contains the gender assertion
type Gender struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Attribution Attribution `json:"attribution"`
}

// Name contains a name assertion
type Name struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Preferred   bool        `json:"preferred"`
	Attribution Attribution `json:"attribution"`
	NameForms   []*NameForm `json:"nameForms"`
}

// NameForm contains a language-specific name form
type NameForm struct {
	Lang     string      `json:"lang"`
	FullText string      `json:"fullText"`
	Parts    []*NamePart `json:"parts"`
}

// NamePart contains types of http://gedcomx.org/Prefix, Given, Surname, and Suffix
type NamePart struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Fact contains information about a fact
type Fact struct {
	ID          string           `json:"id"`
	Type        string           `json:"type"`
	Value       string           `json:"value"`
	Attribution Attribution      `json:"attribution"`
	Date        Date             `json:"date"`
	Place       Place            `json:"place"`
	Qualifiers  []*FactQualifier `json:"qualifiers"`
}

// FactQualifier contains a Name of http://familysearch.org/v1/Event and a Value of false if this is a custom non-event (i.e., fact)
type FactQualifier struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Date contains a date
type Date struct {
	Original   string     `json:"original"`
	Formal     string     `json:"formal"`
	Normalized []*FSValue `json:"normalized"`
}

// Place contains a place
type Place struct {
	Original    string     `json:"original"`
	Description string     `json:"description"`
	Normalized  []*FSValue `json:"normalized"`
}

// GetPersonWithRelationships https://familysearch.org/developers/docs/api/tree/Person_With_Relationships_resource
func (c *Client) GetPersonWithRelationships(pid string) (*PersonWithRelationships, error) {
	u, err := c.GetURL("person-with-relationships-query", nil)
	if err != nil {
		return nil, err
	}
	personWithRelationships := &PersonWithRelationships{}
	err = c.Get(u, map[string]string{"person": pid, "persons": "true"}, nil, personWithRelationships)
	return personWithRelationships, err
}


// GetPersonPortraitURL returns the URL of the person portrait or the empty string
func (c *Client) GetPersonPortraitURL(pid string) (string, error) {
	u, err := c.GetURL("person-portrait-template", map[string]string{"pid": pid})
	if err != nil {
		return "", err
	}
	res, err := c.HTTP("GET", u, nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return "", nil
	} else if res.StatusCode >= 300 && res.StatusCode <= 399 {
		return res.Header.Get("Location"), nil
	}
	return "", fmt.Errorf("Status code %d", res.StatusCode)
}

// GetPreferredParentsURL returns the ID of the preferred parent relationship for the specified user.TreeUserID and personID
// or the empty string if no parents are preferred
func (c *Client) GetPreferredParentsID(tuid, pid string) (string, error) {
	u, err := c.GetURL("preferred-parent-relationship-template", map[string]string{"uid": tuid, "pid": pid})
	if err != nil {
		return "", err
	}
	res, err := c.HTTP("GET", u, nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return "", nil
	} else if res.StatusCode >= 300 && res.StatusCode <= 399 {
		url := strings.Split(res.Header.Get("Location"), "?")[0] // remove query if any
		return url[strings.LastIndex(url,"/")+1:] , nil
	}
	return "", fmt.Errorf("Status code %d", res.StatusCode)
}
