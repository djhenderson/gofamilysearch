package gofamilysearch

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"fmt"
	"errors"
	"io/ioutil"
)

// The Context can be shared among go-routines
type Context struct {
	Environment string
	once 		sync.Once
	templates   map[string]string
}

var apiServer map[string]string = map[string]string{
	"sandbox": "https://sandbox.familysearch.org",
}

func (c *Context) getTemplate(key string, client *http.Client) (string, error) {
	var err error
	if c.templates == nil {
		c.once.Do(func() {
			c.templates, err = readTemplates(apiServer[c.Environment], client)
		})
		if c.templates == nil && err == nil {
			err = errors.New("templates not read")
		}
		if err != nil {
			return "", err
		}
	}
	template, ok := c.templates[key]
	if !ok {
		err = fmt.Errorf("key %s not found", key)
	}
	return template, err
}

type discoveryResponse struct {
	Collections []*discoveryCollection `json:"collections"`
}
type discoveryCollection struct {
  ID string `json:"id"`
  Links map[string]discoveryLink `json:"links"`
}

type discoveryLink struct {
  Template *string `json:"template"`
  Href *string `json:"href"`
}

var templateRegexp = regexp.MustCompile("{\\?[^}]*}")

func readTemplates(host string, client *http.Client) (map[string]string, error) {
	// read discovery url
	req, err := http.NewRequest("GET", host+"/.well-known/collections/tree", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/x-fs-v1+json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Discovery status code %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &discoveryResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	templates := make(map[string]string)
	fsftCollection := func([]*discoveryCollection) *discoveryCollection  {
		for _, coll := range response.Collections {
			if coll.ID == "FSFT" {
				return coll
			}
		}
		return nil
	}(response.Collections)

	for k, v := range fsftCollection.Links {
		var value string
		if v.Href != nil {
			value = *v.Href
		} else {
			value = templateRegexp.ReplaceAllString(*v.Template, "")
		}
		if strings.Index(value, "/") == 0 {
			value = host + value
		}
		templates[k] = value
	}
	return templates, nil
}
