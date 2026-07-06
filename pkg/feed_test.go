package gohumble

import (
	"strings"
	"testing"
	"time"
)

const sampleJSON = `{
	"data": {
		"popular": {"mosaic": [{"products": [{"product_url": "/popular/ignored", "tile_short_name": "Popular"}]}]},
		"books": {
			"mosaic": [
				{"products": [
					{
						"product_url": "/books/cool-bundle",
						"tile_short_name": "Cool Bundle",
						"short_marketing_blurb": "Short blurb",
						"detailed_marketing_blurb": "Detailed blurb",
						"start_date|datetime": "2026-07-01T18:00:00"
					},
					{
						"product_url": "/books/older-bundle",
						"tile_short_name": "Older Bundle",
						"short_marketing_blurb": "Old",
						"detailed_marketing_blurb": "Old detailed",
						"start_date|datetime": "2026-06-15T18:00:00"
					}
				]},
				{"products": [
					{
						"product_url": "/books/cool-bundle",
						"tile_short_name": "Cool Bundle Duplicate",
						"start_date|datetime": "2026-07-01T18:00:00"
					},
					{
						"product_url": "/books/broken-date",
						"tile_short_name": "Broken Date",
						"start_date|datetime": "not-a-date"
					},
					{
						"product_url": "/books/no-date",
						"tile_short_name": "No Date"
					}
				]}
			]
		}
	}
}`

func TestParseProducts(t *testing.T) {
	products, err := parseProducts([]byte(sampleJSON), "books")
	if err != nil {
		t.Fatalf("parseProducts: %v", err)
	}
	// 5 entries in the JSON, one is a duplicate URL -> 4 unique products.
	if len(products) != 4 {
		t.Fatalf("got %d products, want 4", len(products))
	}
	if products[0].TileShortName != "Cool Bundle" {
		t.Errorf("got first product %q, want %q", products[0].TileShortName, "Cool Bundle")
	}
}

func TestParseProductsUnknownCategory(t *testing.T) {
	if _, err := parseProducts([]byte(sampleJSON), "games"); err == nil {
		t.Fatal("expected error for category missing from page data")
	}
}

func TestParseProductsInvalidJSON(t *testing.T) {
	if _, err := parseProducts([]byte("<html>not json</html>"), "books"); err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestCreateFeed(t *testing.T) {
	products, err := parseProducts([]byte(sampleJSON), "books")
	if err != nil {
		t.Fatalf("parseProducts: %v", err)
	}

	feed, err := createFeed(products, "books")
	if err != nil {
		t.Fatalf("createFeed: %v", err)
	}

	// Products without a parseable start date are dropped.
	if len(feed.Items) != 2 {
		t.Fatalf("got %d feed items, want 2", len(feed.Items))
	}
	// Newest bundle first.
	if feed.Items[0].Title != "Cool Bundle" {
		t.Errorf("got first item %q, want %q", feed.Items[0].Title, "Cool Bundle")
	}
	if !feed.Items[0].Created.After(feed.Items[1].Created) {
		t.Error("items are not sorted newest-first")
	}
	// Feed timestamp tracks the newest item, not time.Now(), so reruns
	// without new bundles produce byte-identical output.
	want := time.Date(2026, 7, 1, 18, 0, 0, 0, time.UTC)
	if !feed.Created.Equal(want) {
		t.Errorf("got feed created %v, want %v", feed.Created, want)
	}
	if feed.Title != "Go Humble! RSS Books" {
		t.Errorf("got feed title %q", feed.Title)
	}

	rss, err := feed.ToRss()
	if err != nil {
		t.Fatalf("ToRss: %v", err)
	}
	if !strings.Contains(rss, "https://www.humblebundle.com/books/cool-bundle") {
		t.Error("rss output is missing the product link")
	}
}

func TestCreateFeedRefusesEmptyFeed(t *testing.T) {
	products := []Product{{TileShortName: "No Date", ProductURL: "/books/no-date"}}
	if _, err := createFeed(products, "books"); err == nil {
		t.Fatal("expected error when no product yields a valid feed item")
	}
}
