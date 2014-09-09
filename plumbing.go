package familysearch

import (
	"github.com/rootsdev/familysearch/helpers"
	"encoding/xml"
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func (fs *FamilySearch) getUrl(key string, params map[string]string) (*url.URL, error) {
	segments := regexp.MustCompile("[{}]").Split(fs.Context.getTemplate(key), -1)
	for i, segment := range segments {
		if i % 2 == 1 {
			segments[i] = params[segment]
		}
	}
	// encode the parameters
	// go doesn't have a function to encode path, just a query component, so we convert to a URL for the encoding
	u, err := url.Parse(strings.Join(segments, ""))
	return u, err
}

func (fs *FamilySearch) Get(u url.URL, params map[string]string, headers map[string]string, target interface{}) error {
	helpers.AppendQueryParameters(&u, params)
	body, err := fs.http("GET", u,
					helpers.Extend(map[string]string{"Accept":"application/x-fs-v1+xml"}, headers))
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, target)
}

func (fs *FamilySearch) http(method string, u url.URL, headers map[string]string) ([]byte, error) {
	if fs.AccessToken != "" {
		headers = helpers.Extend(map[string]string{"Authorization":"Bearer "+fs.AccessToken}, headers)
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
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
