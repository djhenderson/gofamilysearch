package gofamilysearch

import "strconv"

type searchMatchResponse struct {
	Entries []*SearchMatchResult `json:"entries"`
}

// SearchMatchResult contains information about a search or match result
type SearchMatchResult struct {
	ID         string             `json:"id"`
	Score      float32            `json:"score"`
	Confidence int                `json:"confidence"`
	Links      map[string]*FSHref `json:"links"`
	Published  int                `json:"published"`
	Title      string             `json:"title"`
	MatchInfo  []*MatchInfo       `json:"matchInfo"`
}

// MatchInfo contains the collection (e.g., https://familysearch.org/platform/collections/records)
// and status (e.g., http://familysearch.org/v1/Pending)
type MatchInfo struct {
	Collection string `json:"collection"`
	Status     string `json:"status"`
}

// GetPersonMatches https://familysearch.org/developers/docs/api/tree/Person_Matches_resource
func (c *Client) GetPersonMatches(pid string, collection string, confidence int, count int) ([]*SearchMatchResult, error) {
	u, err := c.GetURL("person-matches-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	searchMatchResponse := &searchMatchResponse{}
	params := map[string]string{"confidence": strconv.Itoa(confidence), "count": strconv.Itoa(count)}
	if collection != "" {
		params["collection"] = collection
	}
	if err = c.Get(u, params, nil, searchMatchResponse); err != nil {
		return nil, err
	}
	return searchMatchResponse.Entries, nil
}
