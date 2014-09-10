package gofamilysearch

import (
	"encoding/xml"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Environment struct {
	clientId    string
	environment string
	templates   map[string]string
}

var apiServer map[string]string = map[string]string{
	"sandbox": "https://sandbox.familysearch.org",
}

// Create an Environment. The Environment can be shared among go-routines
// Pass in client to allow running on appengine
// Environment makes an http call to read the discovery url
func NewEnvironment(clientId string, environment string, client *http.Client) *Environment {
	return &Environment{
		clientId:    clientId,
		environment: environment,
		templates:   readTemplates(apiServer[environment], client),
	}
}

func (c *Environment) getTemplate(key string) string {
	return c.templates[key]
}

type atomFeed struct {
	Link []atomLink `xml:"link"`
}
type atomLink struct {
	Rel      string `xml:"rel,attr"`
	Href     string `xml:"href,attr"`
	Template string `xml:"template,attr"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readTemplates(host string, client *http.Client) map[string]string {
	// read discovery url
	// NOTE: this endpoint is being deprecated, but the new one is still listed as experimental
	req, err := http.NewRequest("GET", host+"/.well-known/app-meta", nil)
	check(err)
	req.Header.Set("Accept", "application/atom+xml")
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Discovery status code %d", res.StatusCode)
	}
	var discoveryResponse atomFeed
	err = xml.NewDecoder(res.Body).Decode(&discoveryResponse)
	check(err)

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
	return templates
}
