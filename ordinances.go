package gofamilysearch

// GetOrdinanceAccess https://familysearch.org/developers/docs/api/ordinances/Ordinances_resource
func (c *Client) GetOrdinanceAccess() (bool, error) {
	u, err := c.GetURL("ordinances", nil)
	if err != nil {
		return false, err
	}
	res, err := c.HTTP("GET", u, map[string]string{"Accept": "application/x-fs-v1+json"})
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}
