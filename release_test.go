package discogs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRelease_GetRelease(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/249504", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"title": "Never Gonna Give You Up", "id": 249504, "year": 1987}`)
	})

	release, _, err := client.Release.Get(ctx, 249504)
	if err != nil {
		t.Errorf("Release.Get returned error: %v", err)
	}

	expected := &Release{ID: 249504, Title: "Never Gonna Give You Up", Year: 1987}
	if !reflect.DeepEqual(release, expected) {
		t.Errorf("Release.Get\n got=%#v\nwant=%#v", release, expected)
	}
}
