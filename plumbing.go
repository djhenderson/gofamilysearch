package gofamilysearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
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
	retries := 3
	for {
		res, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode >= 200 && res.StatusCode <= 299 {
			return ioutil.ReadAll(res.Body)
		} else if res.StatusCode >= 300 && res.StatusCode <= 399 {
			return []byte(res.Header.Get("Location")), nil
		} else if res.StatusCode == 429 { // throttling response
			time.Sleep(500)
		} else if res.StatusCode >= 500 && method == "GET" && retries > 0 { // possibly-transient error
			retries--
		} else { // error
			return nil, fmt.Errorf("Status code %d", res.StatusCode)
		}
	}
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
	Links map[string]*discoveryLink `json:"links"`
}

type discoveryLink struct {
	Template string `json:"template"`
	Href     string `json:"href"`
}

func (c *Client) readDiscoveryResource(host string) (map[string]string, error) {
	// read discovery url
	u, err := url.Parse(host + "/.well-known/app-meta")
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
	for k, v := range response.Links {
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
