package gofamilysearch

// ChildAndParentsRelationship contains information about a child and parents relationship
type ChildAndParentsRelationship struct {
	ID          string             `json:"id"`
	Links       map[string]*FSHref `json:"links"`
	FatherFacts []*Fact            `json:"fatherFacts"`
	MotherFacts []*Fact            `json:"motherFacts"`
	Father      ResourceRef        `json:"father"`
	Mother      ResourceRef        `json:"mother"`
	Child       ResourceRef        `json:"child"`
}
