package gofamilysearch

import "fmt"

type userResponse struct {
	Users []*User `json:"users"`
}

type User struct {
	Id                string `json:"id"`
	PersonId          string `json:"personId"`
	TreeUserId        string `json:"treeUserId"`
	ContactName       string `json:"contactName"`
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	FamilyName        string `json:"familyName"`
	Gender            string `json:"gender"`
	Email             string `json:"email"`
	PreferredLanguage string `json:"preferredLanguage"`
}

func (c *Client) GetCurrentUser() (*User, error) {
	u, err := c.getUrl("current-user", nil)
	if err != nil {
		return nil, err
	}
	userResponse := &userResponse{}
	err = c.Get(u, nil, nil, userResponse)
	if err != nil {
		return nil, err
	}
	if len(userResponse.Users) != 1 {
		return nil, fmt.Errorf("User not found in response")
	}
	return userResponse.Users[0], nil
}
