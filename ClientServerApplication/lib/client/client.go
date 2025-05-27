package client

import (
	"ClientServerApplication/lib/report"
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	Id     int
	Client *http.Client
}

func (c Client) Send(url string, report *report.Report) (*http.Response, error) {
	content, err := json.Marshal(report)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return resp, nil
}
