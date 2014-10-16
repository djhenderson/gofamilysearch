package gofamilysearch

import "fmt"

type userResponse struct {
	Users []*User `json:"users"`
}

type agentResponse struct {
	Agents []*Agent `json:"agents"`
}

// User contains information about the current user (not to be confused with an Agent)
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

// Agent contains information about another user (not to be confused with User)
type Agent struct {
	ID        string          `json:"id"`
	Names     []*AgentName    `json:"names"`
	Accounts  []*AgentAccount `json:"accounts"`
	Emails    []*AgentEmail   `json:"emails"`
	Phones    []*AgentPhone   `json:"phones"`
	Addresses []*AgentAddress `json:"addresses"`
}

// AgentName contains the agent's name in Value
type AgentName struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// AgentAccount contains the agent's account name
type AgentAccount struct {
	AccountName string `json:"accountName"`
}

// AgentEmail contains the agent's email
type AgentEmail struct {
	Resource string `json:"resource"`
}

// AgentPhone contains the agent's phone
type AgentPhone struct {
	Resource string `json:"resource"`
}

// AgentAddress contains the agent's address
type AgentAddress struct {
	Value string `json:"value"`
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

// GetAgent https://familysearch.org/developers/docs/api/users/Agent_resource
func (c *Client) GetAgent(uid string) (*Agent, error) {
	u, err := c.GetURL("agent-template", map[string]string{"uid": uid})
	if err != nil {
		return nil, err
	}
	agentResponse := &agentResponse{}
	if err = c.Get(u, nil, nil, agentResponse); err != nil {
		return nil, err
	}
	if len(agentResponse.Agents) != 1 {
		return nil, fmt.Errorf("Invalid response")
	}
	return agentResponse.Agents[0], nil
}
