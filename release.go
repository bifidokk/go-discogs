package discogs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const releaseBasePath = "/releases/"

type ReleaseService interface {
	Get(context.Context, int, string) (*Release, *Response, error)
}

type ReleaseServiceOp struct {
	client *Client
}

var _ ReleaseService = &ReleaseServiceOp{}

type Release struct {
	Title             string          `json:"title"`
	ID                int             `json:"id"`
	Artists           []ReleaseArtist `json:"artists"`
	DataQuality       string          `json:"data_quality"`
	Thumb             string          `json:"thumb"`
	Community         []Community     `json:"community"`
	Country           string          `json:"country"`
	Companies         []Company       `json:"companies"`
	DateAdded         string          `json:"date_added"`
	DateChanged       string          `json:"date_changed"`
	EstimatedWeight   int             `json:"estimated_weight"`
	ExtraArtists      []ReleaseArtist `json:"extraartists"`
	FormatQuantity    int             `json:"format_quantity"`
	Formats           []Format        `json:"formats"`
	Genres            []string        `json:"genres"`
	Identifiers       []Identifier    `json:"identifiers"`
	Images            []Image         `json:"images"`
	Labels            []LabelSource   `json:"labels"`
	LowestPrice       float64         `json:"lowest_price"`
	MasterID          int             `json:"master_id"`
	MasterURL         string          `json:"master_url"`
	Notes             string          `json:"notes,omitempty"`
	NumForSale        int             `json:"num_for_sale,omitempty"`
	Released          string          `json:"released"`
	ReleasedFormatted string          `json:"released_formatted"`
	ResourceURL       string          `json:"resource_url"`
	Series            []Series        `json:"series"`
	Status            string          `json:"status"`
	Styles            []string        `json:"styles"`
	Tracklist         []Track         `json:"tracklist"`
	URI               string          `json:"uri"`
	Videos            []Video         `json:"videos"`
	Year              int             `json:"year"`
}

func (rls *ReleaseServiceOp) Get(ctx context.Context, releaseID int, cur string) (*Release, *Response, error) {
	currAbbr, err := currency(cur)
	params := url.Values{}
	params.Set("curr_abbr", currAbbr)

	path := fmt.Sprintf("%s/%d", releaseBasePath, releaseID)

	req, err := rls.client.NewRequest(ctx, http.MethodGet, path, params, nil)
	if err != nil {
		return nil, nil, err
	}

	release := new(Release)
	resp, err := rls.client.Do(ctx, req, release)
	if err != nil {
		return nil, resp, err
	}

	return release, resp, err
}
