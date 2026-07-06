package gohumble

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
)

var categories = []string{"books", "games", "software"}

var httpClient = &http.Client{Timeout: 30 * time.Second}

// Run updates the RSS feed for every category and writes the .rss files into
// outDir. It returns an error if any category fails, so a scheduled run can
// exit non-zero and surface the failure instead of silently going stale.
func Run(outDir string) error {
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
	)
	for _, category := range categories {
		wg.Add(1)
		go func(category string) {
			defer wg.Done()
			if err := updateCategory(category, outDir); err != nil {
				slog.Error("updating category failed", "category", category, "error", err)
				mu.Lock()
				errs = append(errs, fmt.Errorf("%s: %w", category, err))
				mu.Unlock()
			} else {
				slog.Info("feed updated", "category", category)
			}
		}(category)
	}
	wg.Wait()
	return errors.Join(errs...)
}

func updateCategory(category, outDir string) error {
	products, err := fetchProducts(category)
	if err != nil {
		return err
	}

	feed, err := createFeed(products, category)
	if err != nil {
		return err
	}

	return writeFeedToFile(feed, filepath.Join(outDir, category+".rss"))
}

func fetchProducts(category string) ([]Product, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.humblebundle.com/"+category, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gohumble-rss/1.0 (+https://github.com/Feuerlord2/Humble-RSS-Site)")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	data := doc.Find("script#landingPage-json-data").First().Text()
	if data == "" {
		return nil, errors.New("script#landingPage-json-data not found in page")
	}

	return parseProducts([]byte(data), category)
}

func parseProducts(data []byte, category string) ([]Product, error) {
	var page landingPage
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, err
	}

	raw, ok := page.Data[category]
	if !ok {
		return nil, fmt.Errorf("no %q section found in page data", category)
	}

	var cd categoryData
	if err := json.Unmarshal(raw, &cd); err != nil {
		return nil, err
	}
	if len(cd.Mosaic) == 0 {
		return nil, fmt.Errorf("no mosaic data found for %s", category)
	}

	// Collect products across all mosaic sections, deduplicated by URL.
	seen := make(map[string]bool)
	var products []Product
	for _, section := range cd.Mosaic {
		for _, p := range section.Products {
			if p.ProductURL == "" || seen[p.ProductURL] {
				continue
			}
			seen[p.ProductURL] = true
			products = append(products, p)
		}
	}
	return products, nil
}

func createFeed(products []Product, category string) (*feeds.Feed, error) {
	feed := &feeds.Feed{
		Title:       fmt.Sprintf("Go Humble! RSS %s", strings.ToUpper(category[:1])+category[1:]),
		Link:        &feeds.Link{Href: "https://feuerlord2.github.io/Humble-RSS-Site/"},
		Description: fmt.Sprintf("Awesome RSS Feeds about HumbleBundle %s bundles!", category),
		Author:      &feeds.Author{Name: "Daniel Winter", Email: "DanielWinterEmsdetten+rss@gmail.com"},
	}

	for _, product := range products {
		if product.StartDateDatetime == "" {
			slog.Warn("skipping product with empty start date", "product", product.TileShortName)
			continue
		}

		// Humble Bundle emits naive UTC timestamps ("2026-07-01T18:00:00");
		// also accept a zone-suffixed RFC3339 value in case that ever changes.
		dt, err := time.Parse("2006-01-02T15:04:05", product.StartDateDatetime)
		if err != nil {
			dt, err = time.Parse(time.RFC3339, product.StartDateDatetime)
		}
		if err != nil {
			slog.Warn("skipping product with invalid date format",
				"product", product.TileShortName, "date", product.StartDateDatetime)
			continue
		}

		link := "https://www.humblebundle.com" + product.ProductURL
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       product.TileShortName,
			Link:        &feeds.Link{Href: link},
			Id:          link,
			Content:     product.DetailedMarketingBlurb,
			Created:     dt,
			Description: product.ShortMarketingBlurb,
		})
	}

	// Never publish an empty feed: an upstream markup change should fail the
	// run loudly rather than wipe the existing feed file.
	if len(feed.Items) == 0 {
		return nil, fmt.Errorf("no valid products for %s", category)
	}

	// Sort items so that latest bundles are on the top.
	sort.Slice(feed.Items, func(i, j int) bool { return feed.Items[i].Created.After(feed.Items[j].Created) })

	// Stamp the feed with the newest item date instead of time.Now() so the
	// output file only changes when the bundles actually change.
	feed.Created = feed.Items[0].Created

	return feed, nil
}

func writeFeedToFile(feed *feeds.Feed, path string) error {
	rss, err := feed.ToRss()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(rss+"\n"), 0o644)
}
