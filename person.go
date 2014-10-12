package gofamilysearch

type PersonWithRelationships struct {
}

type Person struct {
}

type ChildAndParentsRelationship struct {
}

type CoupleRelationship struct {
}

func (c *Client) GetPersonWithRelationships(pid string) (*PersonWithRelationships, error) {
	u, err := c.GetUrl("person-with-relationships", nil)
	if err != nil {
		return nil, err
	}
	personWithRelationships := &PersonWithRelationships{}
	err = c.Get(u, map[string]string{"person": pid, "persons": "true"}, nil, personWithRelationships)
	return personWithRelationships, err
}
