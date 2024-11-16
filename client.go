package gobalt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
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

func (c *Cobalt) WithHTTPClient(client *http.Client) *Cobalt {
	c.client = client
	return c
}

// Post will return a PostResponse from where the file can be downloaded
// headers are passed as key value pairs. Examples `"API-KEY", "MyApiKey"`
func (c *Cobalt) Post(ctx context.Context, params PostRequest, headers ...string) (*PostResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(params); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.apiBaseURL, buff)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if len(headers)%2 != 0 {
		return nil, fmt.Errorf("odd number of headers params, they must be passed as key value pairs")
	}

	for i := 0; i < len(headers); i += 2 {
		req.Header.Add(headers[i], headers[i+1])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	media := &PostResponse{client: c.client}
	if err := json.NewDecoder(resp.Body).Decode(media); err != nil {
		return nil, err
	}

	if media.Status == ResponseStatusError {
		return nil, CobaltAPIError(*media)
	}

	return media, nil
}

// Stream is a helper utility that will return an io.ReadCloser using the URL from this media object
// The returned io.ReadCloser is the Body of *http.Response and must be closed when you are done with the stream.
// When the m.Status == ResponseStatusPicker it will stream the first item from the m.Picker array.
func (m *PostResponse) Stream(ctx context.Context) (io.ReadCloser, error) {
	if m.Status != ResponseStatusTunnel && m.Status != ResponseStatusRedirect && m.Status != ResponseStatusPicker {
		return nil, fmt.Errorf("unstreamable response type %s", m.Status)
	}

	url := m.URL
	if m.Status == ResponseStatusPicker && len(m.Picker) > 0 {
		url = m.Picker[0].URL
	}
	if len(url) == 0 {
		return nil, fmt.Errorf("url is empty, nothing to stream")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *Cobalt) Get(ctx context.Context, headers ...string) (*GetResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.apiBaseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if len(headers)%2 != 0 {
		return nil, fmt.Errorf("odd number of headers params, they must be passed as key value pairs")
	}

	for i := 0; i < len(headers); i += 2 {
		req.Header.Add(headers[i], headers[i+1])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info := &GetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(info); err != nil {
		return nil, err
	}

	return info, nil
}

const (
	EndpointSession = "session"
)

func (c *Cobalt) Session(ctx context.Context, headers ...string) (*SessionResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.Join(c.apiBaseURL, EndpointSession), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if len(headers)%2 != 0 {
		return nil, fmt.Errorf("odd number of headers params, they must be passed as key value pairs")
	}

	for i := 0; i < len(headers); i += 2 {
		req.Header.Add(headers[i], headers[i+1])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	token := &SessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, err
	}

	if token.Status == ResponseStatusError {
		return nil, fmt.Errorf("%+v", token.ErrorInfo)
	}

	return token, nil
}

// CobalAPIError is just a convenient type to convert Media into an error.
type CobaltAPIError PostResponse

func (err CobaltAPIError) Error() string {
	return fmt.Sprintf("%+v", err.ErrorInfo)
}
