package discogs

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "https://api.discogs.com/"
	userAgent      = "BifidokkDiscogsGoAPIClient/" + libraryVersion
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	return c
}

type ClientOpt func(*Client) error

func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
