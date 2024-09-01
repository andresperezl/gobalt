package gobalt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	CobaltPublicAPI = "https://api.cobalt.tools/api"

	EndpointJSON       = "/json"
	EndpointStream     = "/stream"
	EndpointServerInfo = "/serverInfo"
)

type Cobalt struct {
	client     *http.Client
	apiBaseURL string
}

func NewCobaltWithAPI(apiBaseURL string) *Cobalt {
	return &Cobalt{
		client:     http.DefaultClient,
		apiBaseURL: apiBaseURL,
	}
}

func NewCobaltWithPublicAPI() *Cobalt {
	return &Cobalt{
		client:     http.DefaultClient,
		apiBaseURL: CobaltPublicAPI,
	}
}

func (c *Cobalt) WithHTTPClient(client *http.Client) *Cobalt {
	c.client = client
	return c
}

// Get will return a Response from where the file can be downloaded
func (c *Cobalt) Get(ctx context.Context, params Request) (*Response, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(params); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", c.apiBaseURL, EndpointJSON)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buff)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

// Stream is a helper utility that will return an io.ReadCloser using the URL returned from Get()
// The returned io.ReadCloser is the Body of *http.Response and must be closed when you are done with the stream.
func (c *Cobalt) Stream(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
