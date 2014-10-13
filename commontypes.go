package gofamilysearch

// Attribution contains information about a modification
type Attribution struct {
	Modified      int         `json:"modified"`
	ChangeMessage string      `json:"changeMessage`
	Contributor   ResourceRef `json:"contributor"`
}

// ResourceRef contains information about a resource -- the id and URL
type ResourceRef struct {
	ResourceID string `json:"resourceId"`
	Resource   string `json:"resource"`
}

// A Link contains a single element as far as I know -- href
type Link struct {
	Href string `json:"href"`
}

// NormalizedValue contains a single element as far as I know -- value -- though it may also contain a lang
type NormalizedValue struct {
	Value string `json:"value"`
}
