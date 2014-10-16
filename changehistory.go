package gofamilysearch

import "strconv"

type changesResponse struct {
	Entries []*Change `json:"entries"`
}

// Change contains a change-history entry
type Change struct {
	ID           string               `json:"id"`
	Updated      int                  `json:"updated"`
	Title        string               `json:"title"`
	Links        map[string]*FSHref   `json:"links"`
	Contributors []*ChangeContributor `json:"contributors"`
	ChangeInfo   []*ChangeInfo        `json:"changeInfo"`
}

// ChangeContributor contains information about the user making the change
type ChangeContributor struct {
	Name string `json:"name"`
	URI  string `json:"uri"` // Agent URL
}

// ChangeInfo contains information about the change
type ChangeInfo struct {
	Operation      string      `json:"operation"`
	ObjectType     string      `json:"objectType"`
	ObjectModifier string      `json:"objectModifier"`
	Reason         string      `json:"reason"`
	Resulting      ResourceRef `json:"resulting"`
	Original       ResourceRef `json:"original"`
}

// GetPersonChanges https://familysearch.org/developers/docs/api/tree/Person_Change_History_resource
func (c *Client) GetPersonChanges(pid string, count int, from string) ([]*Change, error) {
	u, err := c.GetURL("person-changes-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}

	var params map[string]string
	if count > 0 {
		params["count"] = strconv.Itoa(count)
	}
	if from != "" {
		params["from"] = from
	}

	changesResponse := &changesResponse{}
	if err = c.Get(u, params, map[string]string{"Accept": "application/x-gedcomx-atom+json"}, changesResponse); err != nil {
		return nil, err
	}
	return changesResponse.Entries, nil
}
