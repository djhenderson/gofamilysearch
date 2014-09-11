package gofamilysearch

import (
	"encoding/xml"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"fmt"
	"errors"
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

type atomFeed struct {
	Link []atomLink `xml:"link"`
}
type atomLink struct {
	Rel      string `xml:"rel,attr"`
	Href     string `xml:"href,attr"`
	Template string `xml:"template,attr"`
}

func readTemplates(host string, client *http.Client) (map[string]string, error) {
	// read discovery url
	// NOTE: this endpoint is being deprecated, but the new one is still listed as experimental
	req, err := http.NewRequest("GET", host+"/.well-known/app-meta", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/atom+xml")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Discovery status code %d", res.StatusCode)
	}
	var discoveryResponse atomFeed
	if err = xml.NewDecoder(res.Body).Decode(&discoveryResponse); err != nil {
		return nil, err
	}

	templates := make(map[string]string)
	re := regexp.MustCompile("{\\?[^}]*}")
	for _, link := range discoveryResponse.Link {
		value := link.Href
		if len(value) == 0 {
			// we will add query parameters later
			value = re.ReplaceAllString(link.Template, "")
		}
		if strings.Index(value, "/") == 0 {
			value = host + value
		}
		templates[link.Rel] = value
	}
	return templates, nil
}
