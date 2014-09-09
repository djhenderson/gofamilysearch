package helpers

import "net/url"

func Extend(target map[string]string, source map[string]string) map[string]string {
	if source != nil {
		for key, value := range source {
			target[key] = value
		}
	}
	return target
}

func AppendQueryParameters(u *url.URL, params map[string]string) {
	if params != nil {
		values := u.Query()
		for key, value := range params {
			values.Set(key, value)
		}
		u.RawQuery = values.Encode()
	}
}
