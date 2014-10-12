package gofamilysearch

import (
	"io/ioutil"
	"net/url"
	"os"
)

func extend(target map[string]string, source map[string]string) map[string]string {
	if source != nil {
		for key, value := range source {
			target[key] = value
		}
	}
	return target
}

func appendQueryParameters(u *url.URL, params map[string]string) {
	if params != nil {
		values := u.Query()
		for key, value := range params {
			values.Set(key, value)
		}
		u.RawQuery = values.Encode()
	}
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
