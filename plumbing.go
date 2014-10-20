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
        "beta": "https://beta.familysearch.org",
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
	res, err := c.HTTP("GET", u, extend(map[string]string{"Accept": "application/x-fs-v1+json"}, headers))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Status code %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(body) == 0 { // an empty response shouldn't return an error; leave target as-is
		return nil
	}
	return json.Unmarshal(body, target)
}

// HTTP is a low-level call. It adds a header for the access token and retries in case of throttling or transient read error.
// It is the caller's responsibility to close the response body
func (c *Client) HTTP(method string, u *url.URL, headers map[string]string) (*http.Response, error) {
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
		// Must use the low-level Transport call here instead of Client
		// because FS likes to use redirect statuses and we don't want to follow redirects automatically
		res, err := c.Transport.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		if res.StatusCode == 429 { // throttling response; wait for awhile and try again
			time.Sleep(500)
		} else if res.StatusCode >= 500 && method == "GET" && retries > 0 { // possibly-transient error
			retries--
		} else {
			return res, nil
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

	// add missing templates
	templates["ordinances"] = host + "/platform/ordinances/ordinances"

	return templates, nil
}
