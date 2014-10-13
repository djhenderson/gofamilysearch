package gofamilysearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var apiServer = map[string]string{
	"sandbox": "https://sandbox.familysearch.org",
}

var urlTemplateRegexp = regexp.MustCompile("[{}]")

// GetURL constructs the url for the key and params from the discovery resource
func (c *Client) GetURL(key string, params map[string]string) (*url.URL, error) {
	template, err := c.getTemplate(key)
	if err != nil {
		return nil, err
	}
	segments := urlTemplateRegexp.Split(template, -1)
	for i, segment := range segments {
		if i%2 == 1 {
			segments[i] = params[segment]
		}
	}
	// encode the parameters
	// go doesn't have a function to encode path, just a query component, so we convert to a URL for the encoding
	u, err := url.Parse(strings.Join(segments, ""))
	return u, err
}

// Get fetches the contents of the URL into the target
func (c *Client) Get(u *url.URL, params map[string]string, headers map[string]string, target interface{}) error {
	appendQueryParameters(u, params)
	body, err := c.HTTP("GET", u, extend(map[string]string{"Accept": "application/x-fs-v1+json"}, headers))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

// HTTP is the low-level call
func (c *Client) HTTP(method string, u *url.URL, headers map[string]string) ([]byte, error) {
	if c.AccessToken != "" {
		headers = extend(map[string]string{"Authorization": "Bearer " + c.AccessToken}, headers)
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status code %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (c *Client) getTemplate(key string) (string, error) {
	var err error
	if c.Context.templates == nil {
		c.Context.once.Do(func() {
			c.Context.templates, err = c.readDiscoveryResource(apiServer[c.Context.Environment])
		})
		if c.Context.templates == nil && err == nil {
			err = fmt.Errorf("templates not read")
		}
		if err != nil {
			return "", err
		}
	}
	template, ok := c.Context.templates[key]
	if !ok {
		err = fmt.Errorf("key %s not found", key)
	}
	return template, err
}

type discoveryResponse struct {
	Collections []*discoveryCollection `json:"collections"`
}
type discoveryCollection struct {
	ID    string                    `json:"id"`
	Links map[string]*discoveryLink `json:"links"`
}

type discoveryLink struct {
	Template string `json:"template"`
	Href     string `json:"href"`
}

func (c *Client) readDiscoveryResource(host string) (map[string]string, error) {
	// read discovery url
	u, err := url.Parse(host + "/platform/collections/tree")
	if err != nil {
		return nil, err
	}
	response := &discoveryResponse{}
	err = c.Get(u, nil, nil, response)
	if err != nil {
		return nil, err
	}

	return generateTemplates(host, response)
}

var templateRegexp = regexp.MustCompile("{\\?[^}]*}")

func generateTemplates(host string, response *discoveryResponse) (map[string]string, error) {
	templates := make(map[string]string)
	fsftCollection := func([]*discoveryCollection) *discoveryCollection {
		for _, coll := range response.Collections {
			if coll.ID == "FSFT" {
				return coll
			}
		}
		return nil
	}(response.Collections)

	for k, v := range fsftCollection.Links {
		var value string
		if v.Href != "" {
			value = v.Href
		} else {
			value = templateRegexp.ReplaceAllString(v.Template, "")
		}
		if strings.Index(value, "/") == 0 {
			value = host + value
		}
		templates[k] = value
	}
	return templates, nil
}
