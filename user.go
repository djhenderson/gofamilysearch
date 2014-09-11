package gofamilysearch

type userResponse struct {
	User User `xml:"user"`
}

type User struct {
	Id                string `xml:"id,attr"`
	PersonId          string `xml:"personId"`
	TreeUserId        string `xml:"treeUserId"`
	ContactName       string `xml:"contactName"`
	DisplayName       string `xml:"displayName"`
	GivenName         string `xml:"givenName"`
	FamilyName        string `xml:"familyName"`
	Gender            string `xml:"gender"`
	Email             string `xml:"email"`
	PreferredLanguage string `xml:"preferredLanguage"`
}

func (c *Client) GetCurrentUser() (*User, error) {
	u, err := c.getUrl("current-user", nil)
	if err != nil {
		return nil, err
	}
	userResponse := new(userResponse)
	err = c.Get(*u, nil, nil, userResponse)
	return &userResponse.User, err
}
