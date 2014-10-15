package gofamilysearch

import (
	"fmt"
	"strconv"
)

type memoriesResponse struct {
	SourceDescriptions []*Memory `json:"sourceDescriptions"`
}

type personMemoryPersonaRefsResponse struct {
	Persons []*evidenceContainer `json:"persons"`
}

type evidenceContainer struct {
	Evidence []*MemoryPersonaRef `json:"evidence"`
}

type memoryPersonasResponse struct {
	Persons []*MemoryPersona `json:"persons"`
}

// Memory contains information about a memory, including the URL (About) and image/icon/thumbnail Links
type Memory struct {
	ID               string              `json:"id"`
	MediaType        string              `json:"mediaType"`
	ResourceType     string              `json:"resourceType"`
	About            string              `json:"about"`
	ArtifactMetadata []*ArtifactMetadata `json:"artifactMetadata"`
	Created          int                 `json:"created"`
	Titles           []*FSValue          `json:"titles"`
	Descriptions     []*FSValue          `json:"descriptions"`
	Links            map[string]*FSHref  `json:"links"`
}

// ArtifactMetadata contains information about an artifact
type ArtifactMetadata struct {
	Filename     string `json:"filename"`
	ArtifactType string `json:"artifactType"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

// MemoryPersona contains information about a memory persona
type MemoryPersona struct {
	ID      string               `json:"id"`
	Media   []*MediaRef          `json:"media"`
	Names   []*MemoryPersonaName `json:"names"`
	Display MemoryPersonaDisplay `json:"display"`
}

// MediaRef contains a link to the memory and qualifiers; Description is the URL of the memory
type MediaRef struct {
	ID          string            `json:"id"`
	Qualifiers  []*MediaQualifier `json:"qualifiers"`
	Description string            `json:"description"`
}

// MediaQualifier contains a Name (e.g., http://gedcomx.org/RectangleRegion) and a Value (e.g., 0, 0, 1, 1)
type MediaQualifier struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MemoryPersonaName contains (possubly multiple) fulltext names
type MemoryPersonaName struct {
	NameForms []*MemoryPersonaNameForm `json:"nameForms"`
}

// MemoryPersonaNameForm contains a fulltext name
type MemoryPersonaNameForm struct {
	FullText string `json:"fullText"`
}

// MemoryPersonaDisplay contains information about the memory persona in ready-to-display form
type MemoryPersonaDisplay struct {
	Name string `json:"name"`
}

// MemoryPersonaRef links persons to memory personas; ResourceID is the memory persona ID
type MemoryPersonaRef struct {
	ID         string `json:"id"`
	Resource   string `json:"resource"`
	ResourceID string `json:"resourceId"`
}

// GetPersonMemories https://familysearch.org/developers/docs/api/tree/Person_Memories_Query_resource
func (c *Client) GetPersonMemories(pid string, start int, count int) ([]*Memory, error) {
	u, err := c.GetURL("person-memories-query", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	memoriesResponse := &memoriesResponse{}
	if err = c.Get(u, map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count)}, nil,
		memoriesResponse); err != nil {
		return nil, err
	}
	return memoriesResponse.SourceDescriptions, nil
}

// GetPersonMemoryPersonaRefs https://familysearch.org/developers/docs/api/tree/Person_Memory_References_resource
func (c *Client) GetPersonMemoryPersonaRefs(pid string) ([]*MemoryPersonaRef, error) {
	u, err := c.GetURL("person-memory-persona-references-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	personMemoryPersonaRefsResponse := &personMemoryPersonaRefsResponse{}
	if err = c.Get(u, nil, nil, personMemoryPersonaRefsResponse); err != nil {
		return nil, err
	}
	if len(personMemoryPersonaRefsResponse.Persons) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return personMemoryPersonaRefsResponse.Persons[0].Evidence, nil
}

// GetMemoryComments https://familysearch.org/developers/docs/api/memories/Memory_Comments_resource
func (c *Client) GetMemoryComments(mid string) ([]*Comment, error) {
	u, err := c.GetURL("memory-comments-template", map[string]string{"mid": mid})
	if err != nil {
		return nil, err
	}
	commentsResponse := &commentsResponse{}
	if err = c.Get(u, nil, nil, commentsResponse); err != nil {
		return nil, err
	}
	if len(commentsResponse.Discussions) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return commentsResponse.Discussions[0].Comments, nil
}

// GetMemoryPersonas https://familysearch.org/developers/docs/api/memories/Memory_Personas_resource
func (c *Client) GetMemoryPersonas(mid string) ([]*MemoryPersona, error) {
	u, err := c.GetURL("memory-personas-template", map[string]string{"mid": mid})
	if err != nil {
		return nil, err
	}
	memoryPersonasResponse := &memoryPersonasResponse{}
	if err = c.Get(u, nil, nil, memoryPersonasResponse); err != nil {
		return nil, err
	}
	return memoryPersonasResponse.Persons, nil
}
