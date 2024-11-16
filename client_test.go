package gobalt

import (
	"context"
	"net/url"
	"testing"
)

var (
	urls = []string{
		"https://x.com/tonystatovci/status/1856853985149227419?t=WuK-zVfde8WTofpdt7UBaQ&s=19",
	}
)

func TestClient(t *testing.T) {
	client := NewCobaltWithAPI("http://localhost:9000/")
	for _, u := range urls {
		pURL, _ := url.Parse(u)
		t.Run(pURL.Host, func(t *testing.T) {
			media, err := client.Post(context.Background(), PostRequest{URL: u})
			if err != nil {
				t.Errorf("failed to fetch media for %s url with error: %v", u, err)
			}

			if len(media.Filename) == 0 {
				t.Error("filename was empty")
			}

			s, err := media.Stream(context.Background())
			if err != nil {
				t.Errorf("failed to stream media with error: %v", err)
				return
			}
			defer s.Close()
		})
	}
}
