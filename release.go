package discogs

import (
	"context"
	"fmt"
	"net/http"
)

const releaseBasePath = "/releases/"

type ReleaseService interface {
	Get(context.Context, int) (*Release, *Response, error)
}

type ReleaseServiceOp struct {
	client *Client
}

var _ ReleaseService = &ReleaseServiceOp{}

type Release struct {
	Title             string   `json:"title"`
	ID                int      `json:"id"`
	ArtistsSort       string   `json:"artists_sort"`
	DataQuality       string   `json:"data_quality"`
	Thumb             string   `json:"thumb"`
	Country           string   `json:"country"`
	DateAdded         string   `json:"date_added"`
	DateChanged       string   `json:"date_changed"`
	EstimatedWeight   int      `json:"estimated_weight"`
	FormatQuantity    int      `json:"format_quantity"`
	Genres            []string `json:"genres"`
	LowestPrice       float64  `json:"lowest_price"`
	MasterID          int      `json:"master_id"`
	MasterURL         string   `json:"master_url"`
	Notes             string   `json:"notes,omitempty"`
	NumForSale        int      `json:"num_for_sale,omitempty"`
	Released          string   `json:"released"`
	ReleasedFormatted string   `json:"released_formatted"`
	ResourceURL       string   `json:"resource_url"`
	Status            string   `json:"status"`
	Styles            []string `json:"styles"`
	URI               string   `json:"uri"`
	Year              int      `json:"year"`
}

func (rls *ReleaseServiceOp) Get(ctx context.Context, releaseID int) (*Release, *Response, error) {
	path := fmt.Sprintf("%s/%d", releaseBasePath, releaseID)

	req, err := rls.client.NewRequest(ctx, http.MethodGet, path, nil)
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
