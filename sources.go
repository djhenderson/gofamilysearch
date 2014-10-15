package gofamilysearch

import (
	"fmt"
	"strings"
)

type sourcesResponse struct {
	Persons                      []*sourceRefsContainer `json:"persons"`
	Relationships                []*sourceRefsContainer `json:"relationships"`
	ChildAndParentsRelationships []*sourceRefsContainer `json:"childAndParentsRelationships"`
	SourceDescriptions           []*SourceDescription   `json:"sourceDescriptions"`
}

type sourceRefsContainer struct {
	Sources []*SourceRef `json:"sources"`
}

// Sources contains SourceRefs and SourceDescriptions attached to a person, couple, or child-and-parents relationship
type Sources struct {
	SourceRefs         []*SourceRef         `json:"sources"`
	SourceDescriptions []*SourceDescription `json:"sourceDescriptions"`
}

// GetSourceDescription returns the SourceDescription for the specified SourceRef.Description
func (sources *Sources) GetSourceDescription(description string) *SourceDescription {
	if strings.HasPrefix(description, "#") {
		description = description[1:]
	}
	for _, sd := range sources.SourceDescriptions {
		if sd.ID == description {
			return sd
		}
	}
	return nil
}

// SourceDescription contains information about a source description
type SourceDescription struct {
	ID          string      `json:"id"`
	About       string      `json:"about"`
	Attribution Attribution `json:"attribution"`
	Citations   []*FSValue  `json:"citations"`
	Titles      []*FSValue  `json:"titles"`
	Notes       []*FSText   `json:"notes"`
}

// SourceRef contains a reference to a SourceDescription; Description is the link to the SourceDescription
type SourceRef struct {
	ID          string      `json:"id"`
	Attribution Attribution `json:"attribution"`
	Description string      `json:"description"`
	Tags        []*Tag      `json:"tags"`
}

// Tag contains a Resource that is http://gedcomx.org/Name, Gender, Birth, Christening, Death, or Burial
type Tag struct {
	Resource string `json:"resource"`
}

// GetPersonSources https://familysearch.org/developers/docs/api/tree/Person_Sources_Query_resource
func (c *Client) GetPersonSources(pid string) (*Sources, error) {
	u, err := c.GetURL("person-sources-query-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	sourcesResponse := &sourcesResponse{}
	if err = c.Get(u, nil, nil, sourcesResponse); err != nil {
		return nil, err
	}
	if len(sourcesResponse.Persons) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return &Sources{
		SourceRefs:         sourcesResponse.Persons[0].Sources,
		SourceDescriptions: sourcesResponse.SourceDescriptions,
	}, nil
}

// GetCoupleSources https://familysearch.org/developers/docs/api/tree/Couple_Relationship_Sources_Query_resource
func (c *Client) GetCoupleSources(crid string) (*Sources, error) {
	u, err := c.GetURL("couple-relationship-sources-query-template", map[string]string{"crid": crid})
	if err != nil {
		return nil, err
	}
	sourcesResponse := &sourcesResponse{}
	if err = c.Get(u, nil, nil, sourcesResponse); err != nil {
		return nil, err
	}
	if len(sourcesResponse.Relationships) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return &Sources{
		SourceRefs:         sourcesResponse.Relationships[0].Sources,
		SourceDescriptions: sourcesResponse.SourceDescriptions,
	}, nil
}

// GetChildAndParentsSources https://familysearch.org/developers/docs/api/tree/Child-and-Parents_Relationship_Sources_Query_resource
func (c *Client) GetChildAndParentsSources(caprid string) (*Sources, error) {
	u, err := c.GetURL("child-and-parents-relationship-sources-template", map[string]string{"caprid": caprid})
	if err != nil {
		return nil, err
	}
	sourcesResponse := &sourcesResponse{}
	if err = c.Get(u, nil, nil, sourcesResponse); err != nil {
		return nil, err
	}
	if len(sourcesResponse.ChildAndParentsRelationships) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return &Sources{
		SourceRefs:         sourcesResponse.ChildAndParentsRelationships[0].Sources,
		SourceDescriptions: sourcesResponse.SourceDescriptions,
	}, nil
}
