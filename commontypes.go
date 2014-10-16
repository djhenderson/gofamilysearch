package gofamilysearch

// Attribution contains information about a modification
type Attribution struct {
	Modified      int         `json:"modified"`
	ChangeMessage string      `json:"changeMessage"`
	Contributor   ResourceRef `json:"contributor"`
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

// FSValue contains a value, and occasionally a lang (e.g., "en-US")
type FSValue struct {
	Value string `json:"value"`
	Lang  string `json:"lang"`
}

// FSText contains a single element as far as I know -- Text
type FSText struct {
	Text string `json:"text"`
}

// FSDescription contains a single element -- Description
type FSDescription struct {
	Description string `json:"description"`
}
