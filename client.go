package gobalt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path"
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
func (c *Cobalt) Get(ctx context.Context, params Request) (*Media, error) {
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

	media := &Media{client: c.client}
	if err := json.NewDecoder(resp.Body).Decode(media); err != nil {
		return nil, err
	}

	return media, nil
}

// ParseFilename will try to extract the filename depending on the type of the m.StatusResponse.
// This is intended to be used with the *http.Response when calling the URL pointedb by m.URL
//
// When m.StatusResponse == StatusResponseRedirect, the filename will be set based on the basename of the URL path
//
// When m.StatusResponse == StatusResponseStream, the filename will be extracted from the Content-Disposition header
//
// All other unsupported methods leave the m.Filename empty
// Errors returned are unexpected, and will be a consenquence of a parsing error.
func (m *Media) ParseFilename(resp *http.Response) error {
	if m.Status == ResponseStatusError || m.Status == ResponseStatusRateLimit {
		return nil
	}

	if m.Status == ResponseStatusRedirect {
		parsedURL, err := url.Parse(m.URL)
		if err != nil {
			return err
		}
		m.filename = path.Base(parsedURL.Path)
		return nil
	}

	if m.Status == ResponseStatusStream {
		cd := resp.Header.Get("Content-Disposition")
		if cd != "" {
			_, params, err := mime.ParseMediaType(cd)
			if err != nil {
				return err
			}
			if filename, ok := params["filename"]; ok {
				m.filename = filename
			}
		}
	}

	return nil
}

// Filename will return the filename associated with this media. ParseFilename must be called first, either directly or indirectly via m.Stream().
// Not doing so will keep the filename empty.
func (m *Media) Filename() string {
	return m.filename
}

// Stream is a helper utility that will return an io.ReadCloser using the URL from this media object
// The returned io.ReadCloser is the Body of *http.Response and must be closed when you are done with the stream.
// Stream will also call ParseFilename, so m.Filename() will be set
func (m *Media) Stream(ctx context.Context) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, m.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := m.ParseFilename(resp); err != nil {
		defer resp.Body.Close()
		return nil, err
	}

	return resp.Body, nil
}
