package gofamilysearch

import "fmt"

type personDiscussionRefsResponse struct {
	Persons []*discussionRefContainer `json:"persons"`
}

type discussionRefContainer struct {
	DiscussionReferences []*DiscussionRef `json:"discussion-references"`
}

type discussionResponse struct {
	Discussions []*Discussion `json:"discussions"`
}

type commentsResponse struct {
	Discussions []*commentContainer `json:"discussions"`
}

type commentContainer struct {
	Comments []*Comment `json:"comments"`
}

// Discussion contains information about a discussion
type Discussion struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	Details          string      `json:"details"`
	Created          int         `json:"created"`
	Modified         int         `json:"modified"`
	NumberOfComments int         `json:"numberOfComments"`
	Contributor      ResourceRef `json:"contributor"`
}

// DiscussionRef contains a reference to a discussion; ResourceID is the discussion ID
type DiscussionRef struct {
	ID               string      `json:"id"`
	ResourceID       string      `json:"resourceId"`
	Resource         string      `json:"resource"`
	Attribution      Attribution `json:"attribution"`
	Details          string      `json:"details"`
	Created          int         `json:"created"`
	Modified         int         `json:"modified"`
	NumberOfComments int         `json:"numberOfComments"`
	Contributor      ResourceRef `json:"contributor"`
}

// Comment contains information about a comment
type Comment struct {
	ID          string      `json:"id"`
	Text        string      `json:"text"`
	Created     int         `json:"created"`
	Contributor ResourceRef `json:"contributor"`
}

// GetPersonDiscussionRefs https://familysearch.org/developers/docs/api/tree/Person_Discussion_References_resource
func (c *Client) GetPersonDiscussionRefs(pid string) ([]*DiscussionRef, error) {
	u, err := c.GetURL("person-discussion-references-template", map[string]string{"pid": pid})
	if err != nil {
		return nil, err
	}
	personDiscussionRefsResponse := &personDiscussionRefsResponse{}
	if err = c.Get(u, nil, nil, personDiscussionRefsResponse); err != nil {
		return nil, err
	}
	if len(personDiscussionRefsResponse.Persons) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return personDiscussionRefsResponse.Persons[0].DiscussionReferences, nil
}

// GetDiscussion https://familysearch.org/developers/docs/api/discussions/Discussion_resource
func (c *Client) GetDiscussion(did string) (*Discussion, error) {
	u, err := c.GetURL("discussion-template", map[string]string{"did": did})
	if err != nil {
		return nil, err
	}
	discussionResponse := &discussionResponse{}
	if err = c.Get(u, nil, nil, discussionResponse); err != nil {
		return nil, err
	}
	if len(discussionResponse.Discussions) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return discussionResponse.Discussions[0], nil
}

// GetDiscussionComments https://familysearch.org/developers/docs/api/discussions/Comments_resource
func (c *Client) GetDiscussionComments(did string) ([]*Comment, error) {
	u, err := c.GetURL("discussion-comments-template", map[string]string{"did": did})
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
