package lastfm

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiBaseURL = "http://ws.audioscrobbler.com/2.0/?"
)

func buildQueryURL(query map[string]string) string {
	parts := make([]string, 0, len(query))
	for key, value := range query {
		parts = append(parts, strings.Join([]string{url.QueryEscape(key), url.QueryEscape(value)}, "="))
	}
	return apiBaseURL + strings.Join(parts, "&")
}

type getter interface {
	Get(url string) (resp *http.Response, err error)
}

type mockServer interface {
	doQuery(params map[string]string) ([]byte, error)
}

// Struct used to access the API servers.
type LastFM struct {
	apiKey string
	getter getter
}

// Create a new LastFM struct.
// The apiKey parameter must be an API key registered with Last.fm.
func New(apiKey string) LastFM {
	return LastFM{apiKey: apiKey, getter: http.DefaultClient}
}

func (lfm LastFM) doQuery(method string, params map[string]string) (body io.ReadCloser, err error) {
	queryParams := make(map[string]string, len(params)+2)
	queryParams["api_key"] = lfm.apiKey
	queryParams["method"] = method
	for key, value := range params {
		queryParams[key] = value
	}

	resp, err := lfm.getter.Get(buildQueryURL(queryParams))
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return
	}
	return resp.Body, err
}

// Used to unwrap XML from inside the <lfm> parent
type lfmStatus struct {
	Status       string       `xml:"status,attr"`
	RecentTracks RecentTracks `xml:"recenttracks"`
	Tasteometer  Tasteometer  `xml:"comparison"`
	TrackInfo    TrackInfo    `xml:"track"`
	TopTags      TopTags      `xml:"toptags"`
	Neighbours   []Neighbour  `xml:"neighbours>user"`
	TopArtists   TopArtists   `xml:"topartists"`
	Error        LastFMError  `xml:"error"`
}

type lfmDate struct {
	Date string `xml:",chardata"`
	UTS  int64  `xml:"uts,attr"`
}
