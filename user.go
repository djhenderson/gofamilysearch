package gofamilysearch

import "fmt"

type userResponse struct {
	Users []*User `json:"users"`
}

// User contains information about a user (not to be confused with an Agent)
type User struct {
	ID                string `json:"id"`
	PersonID          string `json:"personId"`
	TreeUserID        string `json:"treeUserId"`
	ContactName       string `json:"contactName"`
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	FamilyName        string `json:"familyName"`
	Gender            string `json:"gender"`
	Email             string `json:"email"`
	PreferredLanguage string `json:"preferredLanguage"`
}

// GetCurrentUser https://familysearch.org/developers/docs/api/users/Current_User_resource
func (c *Client) GetCurrentUser() (*User, error) {
	u, err := c.GetURL("current-user", nil)
	if err != nil {
		return nil, err
	}
	userResponse := &userResponse{}
	if err = c.Get(u, nil, nil, userResponse); err != nil {
		return nil, err
	}
	if len(userResponse.Users) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return userResponse.Users[0], nil
}
