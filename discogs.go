package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "https://api.discogs.com/"
	userAgent      = "discogs-go/" + libraryVersion
	mediaType      = "json"
)

type Client struct {
	client      *http.Client
	BaseURL     *url.URL
	UserAgent   string
	AccessToken string

	Release ReleaseService
}

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func NewFromToken(token string) (*Client, error) {
	c, err := New(nil, SetUserToken(token))

	return c, err
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Release = &ReleaseServiceOp{client: c}

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

func SetUserToken(token string) ClientOpt {
	return func(c *Client) error {
		c.AccessToken = token
		return nil
	}
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, params url.Values, body interface{}) (*http.Request, error) {
	if params != nil {
		urlStr = urlStr + "?" + params.Encode()
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Authorization", fmt.Sprintf("Discogs token=%s", c.AccessToken))
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}

	return &response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := DoRequestWithClient(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

func DoRequestWithClient(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}

	return errorResponse
}
