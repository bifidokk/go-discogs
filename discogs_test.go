package discogs

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	ctx    = context.TODO()
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	testClientDefaults(t, c)
}

func TestNewFromToken(t *testing.T) {
	c := NewFromToken("my-token")
	testClientDefaults(t, c)
}

func TestNew(t *testing.T) {
	c, err := New(nil)

	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	testClientDefaults(t, c)
}

func TestCustomUserAgent(t *testing.T) {
	ua := "test-user-agent/1.0.0"
	c, err := New(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := c.UserAgent; got != expected {
		t.Errorf("New() UserAgent = %s; expected %s", got, expected)
	}
}

func TestCustomBaseURL(t *testing.T) {
	baseURL := "https://custom.discogs.com/"
	c, err := New(nil, SetBaseURL(baseURL))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := baseURL
	if got := c.BaseURL.String(); got != baseURL {
		t.Errorf("New() BaseURL = %s; expected %s", got, expected)
	}
}

func TestInvalidBaseURL(t *testing.T) {
	baseURL := ":"
	_, err := New(nil, SetBaseURL(baseURL))

	testURLParseError(t, err)
}

func testClientDefaults(t *testing.T, c *Client) {
	testClientDefaultBaseURL(t, c)
	testClientDefaultUserAgent(t, c)
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.BaseURL == nil || c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL, defaultBaseURL)
	}
}

func testClientDefaultUserAgent(t *testing.T, c *Client) {
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &AddToWantlistRequest{}, `{}`+"\n"
	req, _ := c.NewRequest(ctx, http.MethodGet, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)
	body := new(foo)
	_, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	expected := &foo{"a"}
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}
