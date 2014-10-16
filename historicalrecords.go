package gofamilysearch

import "net/url"

// HistoricalRecord contains information about a historical record
type HistoricalRecord struct {
	ID                 string                               `json:"id"`
	Description        string                               `json:"description"`
	SourceDescriptions []*HistoricalRecordSourceDescription `json:"sourceDescriptions"`
	Relationships      []*HistoricalRecordRelationship      `json:"relationships"`
	Persons            []*HistoricalRecordPerson            `json:"persons"`
	Fields             []*HistoricalRecordField             `json:"fields"`
}

// HistoricalRecordSourceDescription contains a SourceDescription for historical records
type HistoricalRecordSourceDescription struct {
	ID           string                      `json:"id"`
	About        string                      `json:"about"`
	ResourceType string                      `json:"resourceType"`
	Created      int                         `json:"created"`
	Modified     int                         `json:"modified"`
	Citations    []*FSValue                  `json:"citations"`
	Titles       []*FSValue                  `json:"titles"`
	Identifiers  map[string][]string         `json:"identifiers"`
	ComponentOf  FSDescription               `json:"componentOf"`
	Sources      []*FSDescription            `json:"sources"`
	Coverage     []*HistoricalRecordCoverage `json:"coverage"`
}

// HistoricalRecordCoverage contains coverage information
type HistoricalRecordCoverage struct {
	RecordType string                           `json:"recordType"`
	Temporal   HistoricalRecordTemporalCoverage `json:"temporal"`
	Spatial    HistoricalRecordSpatialCoverage  `json:"spatial"`
}

// HistoricalRecordTemporalCoverage contains time coverage
type HistoricalRecordTemporalCoverage struct {
	Formal   string `json:"formal"`
	Original string `json:"original"`
}

// HistoricalRecordSpatialCoverage contains places coverage
type HistoricalRecordSpatialCoverage struct {
	Original string `json:"original"`
}

// HistoricalRecordRelationship contains a Relationship for historical records
type HistoricalRecordRelationship struct {
	ID      string                  `json:"id"`
	Type    string                  `json:"type"`
	Facts   []*HistoricalRecordFact `json:"facts"`
	Person1 ResourceRef             `json:"person1"`
	Person2 ResourceRef             `json:"person2"`
}

// HistoricalRecordPerson contains a Person for historical records
type HistoricalRecordPerson struct {
	ID          string                   `json:"id"`
	Extracted   bool                     `json:"extracted"`
	Principal   bool                     `json:"principal"`
	Gender      HistoricalRecordGender   `json:"gender"`
	Names       []*HistoricalRecordName  `json:"names"`
	Facts       []*HistoricalRecordFact  `json:"facts"`
	Links       map[string]*FSHref       `json:"links"`
	Identifiers map[string][]string      `json:"identifiers"`
	Fields      []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordGender adds Fields to Gender
type HistoricalRecordGender struct {
	Type   string                   `json:"type"`
	Fields []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordName contains a Name for historical records
type HistoricalRecordName struct {
	Type      string                      `json:"type"`
	NameForms []*HistoricalRecordNameForm `json:"nameForms"`
}

// HistoricalRecordNameForm contains a NameForm for historical records
type HistoricalRecordNameForm struct {
	FullText string                      `json:"fullText"`
	Parts    []*HistoricalRecordNamePart `json:"parts"`
	Fields   []*HistoricalRecordField    `json:"fields"`
}

// HistoricalRecordNamePart adds Fields to NamePart
type HistoricalRecordNamePart struct {
	NamePart
	Fields []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordFact adds Primary and Fields to Fact
type HistoricalRecordFact struct {
	Type    string                   `json:"type"`
	Value   string                   `json:"value"`
	Date    HistoricalRecordDate     `json:"date"`
	Place   HistoricalRecordPlace    `json:"place"`
	Primary bool                     `json:"primary"`
	Fields  []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordDate adds Fields to Date
type HistoricalRecordDate struct {
	Date
	Fields []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordPlace adds Fields to Place
type HistoricalRecordPlace struct {
	Place
	Fields []*HistoricalRecordField `json:"fields"`
}

// HistoricalRecordField contains information about a field
type HistoricalRecordField struct {
	Values []*HistoricalRecordFieldValue `json:"values"`
	Type   string                        `json:"type"`
}

// HistoricalRecordFieldValue contains a field value
type HistoricalRecordFieldValue struct {
	Text     string `json:"text"`
	LabelID  string `json:"labelId"`
	Type     string `json:"type"`
	Resource string `json:"resource"`
}

// RecordCollection describes a collection of (historical) records
type RecordCollection struct {
	SourceDescriptions []*RecordCollectionSourceDescription `json:"sourceDescriptions"`
	Collections        []*RecordCollectionCollection        `json:"collections"`
	RecordDescriptors  []*RecordCollectionRecordDescriptor  `json:"recordDescriptors"`
	Description        string                               `json:"description"`
}

// RecordCollectionSourceDescription contains a SourceDescription for a collection
type RecordCollectionSourceDescription struct {
	ID           string                      `json:"id"`
	Citations    []*FSValue                  `json:"citations"`
	About        string                      `json:"about"`
	ComponentOf  FSDescription               `json:"componentOf"`
	Titles       []*FSValue                  `json:"titles"`
	ResourceType string                      `json:"resourceType"`
	Rights       []string                    `json:"rights"`
	Descriptions []*FSValue                  `json:"descriptions"`
	Identifiers  map[string][]string         `json:"identifiers"`
	Coverage     []*HistoricalRecordCoverage `json:"coverage"`
}

// RecordCollectionCollection contains information about the collection
type RecordCollectionCollection struct {
	Lang    string                     `json:"lang"`
	Content []*RecordCollectionContent `json:"content"`
	Title   string                     `json:"title"`
	Size    int                        `json:"size"`
}

// RecordCollectionContent contains information about the collection type and extent
type RecordCollectionContent struct {
	ResourceType string  `json:"resourceType"`
	Count        int     `json:"count"`
	Completeness float32 `json:"completeness"`
}

// RecordCollectionRecordDescriptor contains information for displaying record collection fields
type RecordCollectionRecordDescriptor struct {
	ID     string                   `json:"id"`
	Fields []*RecordCollectionField `json:"fields"`
}

// RecordCollectionField contains a list of values
type RecordCollectionField struct {
	Values []*RecordCollectionFieldValue `json:"values"`
}

// RecordCollectionFieldValue contains language-specific labels for the field
type RecordCollectionFieldValue struct {
	LabelID string     `json:"labelId"`
	Type    string     `json:"type"`
	Labels  []*FSValue `json:"labels"`
}

// GetHistoricalRecord reads the historical record with the specified id
func (c *Client) GetHistoricalRecord(id string) (*HistoricalRecord, error) {
	u, err := url.Parse(id)
	if err != nil {
		return nil, err
	}
	historicalRecord := &HistoricalRecord{}
	if err = c.Get(u, nil, nil, historicalRecord); err != nil {
		return nil, err
	}
	return historicalRecord, nil
}

// GetRecordCollection reads the record collection with the specified id
func (c *Client) GetRecordCollection(id string) (*RecordCollection, error) {
	u, err := url.Parse(id)
	if err != nil {
		return nil, err
	}
	recordCollection := &RecordCollection{}
	if err = c.Get(u, nil, nil, recordCollection); err != nil {
		return nil, err
	}
	return recordCollection, nil
}
