package gallery

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var (
	headers = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json;api-version=3.0-preview.1",
		"User-Agent":   "offvsix (Go)",
	}
)

func newRequest(extensionID string) (*http.Request, error) {
	var (
		scheme   = "https"
		host     = "marketplace.visualstudio.com"
		path     = "/_apis/public/gallery/extensionquery"
		endpoint = scheme + "://" + host + path
	)
	body, err := setupBody(extensionID)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func setupBody(extensionID string) (io.Reader, error) {
	criteria := []Criteria{{FilterType: 7, Value: extensionID}}
	filters := []Filter{{Criteria: criteria}}
	body := Request{Filters: filters, Flags: 914}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(payload), nil
}
