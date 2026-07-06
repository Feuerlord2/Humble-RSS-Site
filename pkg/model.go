package gohumble

import "encoding/json"

// landingPage matches the JSON embedded in the script#landingPage-json-data
// tag on humblebundle.com category pages. The category key under "data"
// ("books", "games", "software") mirrors the page URL, so it is looked up
// dynamically instead of hardcoding one struct per category.
type landingPage struct {
	Data map[string]json.RawMessage `json:"data"`
}

type categoryData struct {
	Mosaic []struct {
		Products []Product `json:"products"`
	} `json:"mosaic"`
}

// Product holds the subset of tile fields the RSS feed needs; unknown JSON
// fields are ignored by encoding/json.
type Product struct {
	ProductURL             string `json:"product_url"`
	TileShortName          string `json:"tile_short_name"`
	ShortMarketingBlurb    string `json:"short_marketing_blurb"`
	DetailedMarketingBlurb string `json:"detailed_marketing_blurb"`
	StartDateDatetime      string `json:"start_date|datetime"`
}
