package gofamilysearch

// Attribution contains information about a modification
type Attribution struct {
	Modified      int         `json:"modified"`
	ChangeMessage string      `json:"changeMessage"`
	Contributor   ResourceRef `json:"contributor"`
}

type commentsResponse struct {
	Discussions []*commentContainer `json:"discussions"`
}

type commentContainer struct {
	Comments []*Comment `json:"comments"`
}

// Comment contains information about a comment
type Comment struct {
	ID          string      `json:"id"`
	Text        string      `json:"text"`
	Created     int         `json:"created"`
	Contributor ResourceRef `json:"contributor"`
}

// ResourceRef contains information about a resource -- the ID and URL
type ResourceRef struct {
	ResourceID string `json:"resourceId"`
	Resource   string `json:"resource"`
}

// A FSHref contains a single element as far as I know -- Href
type FSHref struct {
	Href string `json:"href"`
}

// FSValue contains a single element as far as I know -- Value
type FSValue struct {
	Value string `json:"value"`
}

// FSText contains a single element as far as I know -- Text
type FSText struct {
	Text string `json:"text"`
}
