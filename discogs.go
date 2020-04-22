package discogs

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "https://api.discogs.com/"
	userAgent      = "discogs-go/" + libraryVersion
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

func SetBaseURL(baseURL string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}
