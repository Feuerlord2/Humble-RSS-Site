package gohumble

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	log "github.com/sirupsen/logrus"
)

func Run() {
	wg := sync.WaitGroup{}
	for _, category := range []string{"books", "games", "software"} {
		wg.Add(1)
		go updateCategory(&wg, category)
	}
	wg.Wait()
}

func createFeed(products []Product, category string) feeds.Feed {
	feed := feeds.Feed{
		Title:       fmt.Sprintf("Go Humble! RSS %s", strings.ToUpper(category[:1])+category[1:]),
		Link:        &feeds.Link{Href: "https://feuerlord2.github.io/Humble-RSS-Site/"},
		Description: fmt.Sprintf("Awesome RSS Feeds about HumbleBundle %s bundles!", category),
		Author:      &feeds.Author{Name: "Daniel Winter", Email: "DanielWinterEmsdetten+rss@gmail.com"},
		Created:     time.Now(),
	}

	validProducts := make([]*feeds.Item, 0, len(products))

	for _, product := range products {
		// Skip products without a valid date
		if product.StartDateDatetime == "" {
			log.WithField("product", product.TileShortName).Warn("skipping product with empty StartDateDatetime")
			continue
		}

		// Need to add a Z in order to make it RFC3339 parseable.
		dt, err := time.Parse(time.RFC3339, product.StartDateDatetime+"Z")
		if err != nil {
			log.WithFields(log.Fields{
				"product": product.TileShortName,
				"date":    product.StartDateDatetime,
			}).Warn("skipping product with invalid date format")
			continue
		}

		validProducts = append(validProducts, &feeds.Item{
			Title:       product.TileShortName,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://humblebundle.com%s", product.ProductURL)},
			Content:     product.DetailedMarketingBlurb,
			Created:     dt,
			Description: product.ShortMarketingBlurb,
		})
	}

	feed.Items = validProducts

	// Sort items so that latest bundles are on the top.
	sort.Slice(feed.Items, func(i, j int) bool { return feed.Items[i].Created.After(feed.Items[j].Created) })

	return feed
}

func parseProducts(data []byte, category string) ([]Product, error) {
	switch category {
	case "books":
		var books BooksData
		err := json.Unmarshal(data, &books)
		if err != nil {
			return nil, err
		}
		if len(books.Data.Books.Mosaic) == 0 {
			return nil, fmt.Errorf("no mosaic data found for %s", category)
		}
		return books.Data.Books.Mosaic[0].Products, nil
	case "games":
		var games GamesData
		err := json.Unmarshal(data, &games)
		if err != nil {
			return nil, err
		}
		if len(games.Data.Games.Mosaic) == 0 {
			return nil, fmt.Errorf("no mosaic data found for %s", category)
		}
		return games.Data.Games.Mosaic[0].Products, nil
	case "software":
		var software SoftwareData
		err := json.Unmarshal(data, &software)
		if err != nil {
			return nil, err
		}
		if len(software.Data.Software.Mosaic) == 0 {
			return nil, fmt.Errorf("no mosaic data found for %s", category)
		}
		return software.Data.Software.Mosaic[0].Products, nil
	default:
		return nil, fmt.Errorf("unknown category %s", category)
	}
}

func updateCategory(wg *sync.WaitGroup, category string) {
	defer wg.Done()

	resp, err := http.Get(fmt.Sprintf("https://www.humblebundle.com/%s", category))
	if err != nil {
		log.WithField("category", category).Error(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"category": category,
			"status":   resp.StatusCode,
		}).Error("unexpected HTTP status")
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.WithField("category", category).Error(err)
		return
	}

	doc.Find("script#landingPage-json-data").Each(func(idx int, s *goquery.Selection) {
		node := s.Nodes[0]
		data := node.FirstChild.Data

		products, err := parseProducts([]byte(data), category)
		if err != nil {
			log.WithField("step", "parsing").Error(err)
			return
		}

		feed := createFeed(products, category)

		if err := writeFeedToFile(feed, category); err != nil {
			log.WithField("step", "writing").Error(err)
		}
	})
}

func writeFeedToFile(feed feeds.Feed, category string) error {
	f, err := os.OpenFile(
		fmt.Sprintf("%s.rss", category),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	rss, err := feed.ToRss()
	if err != nil {
		return err
	}

	if _, err := w.WriteString(rss); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
