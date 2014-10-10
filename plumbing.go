package gofamilysearch

import (
	"encoding/json"
	"fmt"
	"github.com/rootsdev/gofamilysearch/helpers"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var urlTemplateRegexp = regexp.MustCompile("[{}]")

func (c *Client) getUrl(key string, params map[string]string) (*url.URL, error) {
	template, err := c.Context.getTemplate(key, c.HttpClient)
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

func (c *Client) Get(u *url.URL, params map[string]string, headers map[string]string, target interface{}) error {
	helpers.AppendQueryParameters(u, params)
	body, err := c.http("GET", u, helpers.Extend(map[string]string{"Accept": "application/x-fs-v1+json"}, headers))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func (c *Client) http(method string, u *url.URL, headers map[string]string) ([]byte, error) {
	if c.AccessToken != "" {
		headers = helpers.Extend(map[string]string{"Authorization": "Bearer " + c.AccessToken}, headers)
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := c.HttpClient.Do(req)
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
