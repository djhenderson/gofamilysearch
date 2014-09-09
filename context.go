package familysearch

import (
	"log"
	"encoding/xml"
	"net/http"
	"regexp"
)

type Context struct {
	clientId string
	environment string
	templates map[string]string
}

type atomFeed struct {
	Link []atomLink `xml:"link"`
}
type atomLink struct {
	Rel string      `xml:"rel,attr"`
	Href string 	`xml:"href,attr"`
	Template string `xml:"template,attr"`
}

var apiServer map[string]string = map[string]string {
	"sandbox" : "https://sandbox.familysearch.org",
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Create a Context. The Context can be shared among go-routines
func NewContext(clientId string, environment string) *Context {
	// read discovery url
	client := &http.Client{}
	// NOTE: this endpoint is being deprecated, but the new one is still listed as experimental
	req, err := http.NewRequest("GET", apiServer[environment]+"/.well-known/app-meta", nil)
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
		templates[link.Rel] = value
	}

	return &Context{
		clientId : clientId,
		environment : environment,
		templates : templates,
	}
}

func (c *Context) getTemplate(key string) string {
	return c.templates[key]
}

func (c *Context) getApiHost() string {
	return apiServer[c.environment]
}
