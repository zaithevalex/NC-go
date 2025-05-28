package client

import (
	"ClientServerApplication/lib/report"
	"bytes"
	"encoding/json"
	"net/http"
)

const contentType = "application/json"

// The Client is an HTTP client with client identifier.
type Client struct {
	Client *http.Client
	Id     int
}

// Send sends report to the specified url.
func (c Client) Send(url string, report *report.Report) (*http.Response, error) {
	content, err := json.Marshal(report)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Post(url, contentType, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return resp, nil
}
