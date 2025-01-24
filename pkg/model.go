package gohumble

// BooksData has been autogenerated based on the JSON content using https://transform.tools/json-to-go.
type BooksData struct {
	UserOptions struct {
		IsLoggedIn bool   `json:"is_logged_in"`
		LogoutURL  string `json:"logout_url"`
	} `json:"userOptions"`
	ShowSingleSignOn           bool `json:"showSingleSignOn"`
	Debug                      bool `json:"debug"`
	IPInChina                  bool `json:"ipInChina"`
	IsEuCountry                bool `json:"isEuCountry"`
	BaseSubscriptionPriceMoney struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"baseSubscriptionPrice|money"`
	BannerOptions interface{} `json:"bannerOptions"`
	ExchangeRates struct {
		USDDecimal float64 `json:"USD|decimal"`
		AUDDecimal float64 `json:"AUD|decimal"`
		CHFDecimal float64 `json:"CHF|decimal"`
		IDRDecimal float64 `json:"IDR|decimal"`
		KRWDecimal float64 `json:"KRW|decimal"`
		BGNDecimal float64 `json:"BGN|decimal"`
		CNYDecimal float64 `json:"CNY|decimal"`
		ISKDecimal float64 `json:"ISK|decimal"`
		ILSDecimal float64 `json:"ILS|decimal"`
		GBPDecimal float64 `json:"GBP|decimal"`
		NZDDecimal float64 `json:"NZD|decimal"`
		DKKDecimal float64 `json:"DKK|decimal"`
		CADDecimal float64 `json:"CAD|decimal"`
		TRYDecimal float64 `json:"TRY|decimal"`
		HUFDecimal float64 `json:"HUF|decimal"`
		PHPDecimal float64 `json:"PHP|decimal"`
		RONDecimal float64 `json:"RON|decimal"`
		NOKDecimal float64 `json:"NOK|decimal"`
		RUBDecimal float64 `json:"RUB|decimal"`
		ZARDecimal float64 `json:"ZAR|decimal"`
		MYRDecimal float64 `json:"MYR|decimal"`
		INRDecimal float64 `json:"INR|decimal"`
		THBDecimal float64 `json:"THB|decimal"`
		MXNDecimal float64 `json:"MXN|decimal"`
		CZKDecimal float64 `json:"CZK|decimal"`
		BRLDecimal float64 `json:"BRL|decimal"`
		JPYDecimal float64 `json:"JPY|decimal"`
		PLNDecimal float64 `json:"PLN|decimal"`
		EURDecimal float64 `json:"EUR|decimal"`
		SEKDecimal float64 `json:"SEK|decimal"`
		SGDDecimal float64 `json:"SGD|decimal"`
		HKDDecimal float64 `json:"HKD|decimal"`
	} `json:"exchangeRates"`
	CsrfTokenInput string `json:"csrfTokenInput"`
	CsrfToken      struct {
	} `json:"csrfToken"`
	Data struct {
		Popular struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string   `json:"section_url"`
				SectionType string   `json:"section_type"`
				Layouts     []string `json:"layouts"`
				Products    []struct {
					TileLogoInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_logo_information"`
					MachineName      string `json:"machine_name"`
					HighResTileImage string `json:"high_res_tile_image"`
					DisableHeroTile  bool   `json:"disable_hero_tile"`
					MarketingBlurb   string `json:"marketing_blurb"`
					HoverTitle       string `json:"hover_title"`
					ProductURL       string `json:"product_url"`
					TileImage        string `json:"tile_image"`
					Category         string `json:"category"`
					HeroHighlights   []struct {
						Heading string `json:"heading"`
						Tooltip string `json:"tooltip"`
					} `json:"hero_highlights"`
					HoverHighlights             []string `json:"hover_highlights"`
					Author                      string   `json:"author"`
					FallbackStoreSaleLogo       string   `json:"fallback_store_sale_logo"`
					HighResTileImageInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"high_res_tile_image_information"`
					SupportsPartners       bool    `json:"supports_partners"`
					DetailedMarketingBlurb string  `json:"detailed_marketing_blurb"`
					TileLogo               string  `json:"tile_logo"`
					TileShortName          string  `json:"tile_short_name"`
					StartDateDatetime      string  `json:"start_date|datetime"`
					EndDateDatetime        string  `json:"end_date|datetime"`
					TileStamp              string  `json:"tile_stamp"`
					BundlesSoldDecimal     float64 `json:"bundles_sold|decimal"`
					TileName               string  `json:"tile_name"`
					ShortMarketingBlurb    string  `json:"short_marketing_blurb"`
					TileImageInformation   struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_image_information"`
					Type       string   `json:"type"`
					Highlights []string `json:"highlights"`
				} `json:"products"`
			} `json:"mosaic"`
		} `json:"popular"`
		Books struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string    `json:"section_url"`
				SectionType string    `json:"section_type"`
				Layouts     []string  `json:"layouts"`
				Products    []Product `json:"products"`
			} `json:"mosaic"`
		} `json:"books"`
	} `json:"data"`
	UserMaxRewardAmount int    `json:"user_max_reward_amount"`
	CsrfFormKey         string `json:"csrfFormKey"`
}

// GamesData has been autogenerated based on the JSON content using https://transform.tools/json-to-go.
type GamesData struct {
	UserOptions struct {
		IsLoggedIn bool   `json:"is_logged_in"`
		LogoutURL  string `json:"logout_url"`
	} `json:"userOptions"`
	ShowSingleSignOn           bool `json:"showSingleSignOn"`
	Debug                      bool `json:"debug"`
	IPInChina                  bool `json:"ipInChina"`
	IsEuCountry                bool `json:"isEuCountry"`
	BaseSubscriptionPriceMoney struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"baseSubscriptionPrice|money"`
	BannerOptions interface{} `json:"bannerOptions"`
	ExchangeRates struct {
		USDDecimal float64 `json:"USD|decimal"`
		AUDDecimal float64 `json:"AUD|decimal"`
		CHFDecimal float64 `json:"CHF|decimal"`
		IDRDecimal float64 `json:"IDR|decimal"`
		KRWDecimal float64 `json:"KRW|decimal"`
		BGNDecimal float64 `json:"BGN|decimal"`
		CNYDecimal float64 `json:"CNY|decimal"`
		ISKDecimal float64 `json:"ISK|decimal"`
		ILSDecimal float64 `json:"ILS|decimal"`
		GBPDecimal float64 `json:"GBP|decimal"`
		NZDDecimal float64 `json:"NZD|decimal"`
		DKKDecimal float64 `json:"DKK|decimal"`
		CADDecimal float64 `json:"CAD|decimal"`
		TRYDecimal float64 `json:"TRY|decimal"`
		HUFDecimal float64 `json:"HUF|decimal"`
		PHPDecimal float64 `json:"PHP|decimal"`
		RONDecimal float64 `json:"RON|decimal"`
		NOKDecimal float64 `json:"NOK|decimal"`
		RUBDecimal float64 `json:"RUB|decimal"`
		ZARDecimal float64 `json:"ZAR|decimal"`
		MYRDecimal float64 `json:"MYR|decimal"`
		INRDecimal float64 `json:"INR|decimal"`
		THBDecimal float64 `json:"THB|decimal"`
		MXNDecimal float64 `json:"MXN|decimal"`
		CZKDecimal float64 `json:"CZK|decimal"`
		BRLDecimal float64 `json:"BRL|decimal"`
		JPYDecimal float64 `json:"JPY|decimal"`
		PLNDecimal float64 `json:"PLN|decimal"`
		EURDecimal float64 `json:"EUR|decimal"`
		SEKDecimal float64 `json:"SEK|decimal"`
		SGDDecimal float64 `json:"SGD|decimal"`
		HKDDecimal float64 `json:"HKD|decimal"`
	} `json:"exchangeRates"`
	CsrfTokenInput string `json:"csrfTokenInput"`
	CsrfToken      struct {
	} `json:"csrfToken"`
	Data struct {
		Popular struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string   `json:"section_url"`
				SectionType string   `json:"section_type"`
				Layouts     []string `json:"layouts"`
				Products    []struct {
					TileLogoInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_logo_information"`
					MachineName      string `json:"machine_name"`
					HighResTileImage string `json:"high_res_tile_image"`
					DisableHeroTile  bool   `json:"disable_hero_tile"`
					MarketingBlurb   string `json:"marketing_blurb"`
					HoverTitle       string `json:"hover_title"`
					ProductURL       string `json:"product_url"`
					TileImage        string `json:"tile_image"`
					Category         string `json:"category"`
					HeroHighlights   []struct {
						Heading string `json:"heading"`
						Tooltip string `json:"tooltip"`
					} `json:"hero_highlights"`
					HoverHighlights             []string `json:"hover_highlights"`
					Author                      string   `json:"author"`
					FallbackStoreSaleLogo       string   `json:"fallback_store_sale_logo"`
					HighResTileImageInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"high_res_tile_image_information"`
					SupportsPartners       bool    `json:"supports_partners"`
					DetailedMarketingBlurb string  `json:"detailed_marketing_blurb"`
					TileLogo               string  `json:"tile_logo"`
					TileShortName          string  `json:"tile_short_name"`
					StartDateDatetime      string  `json:"start_date|datetime"`
					EndDateDatetime        string  `json:"end_date|datetime"`
					TileStamp              string  `json:"tile_stamp"`
					BundlesSoldDecimal     float64 `json:"bundles_sold|decimal"`
					TileName               string  `json:"tile_name"`
					ShortMarketingBlurb    string  `json:"short_marketing_blurb"`
					TileImageInformation   struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_image_information"`
					Type       string   `json:"type"`
					Highlights []string `json:"highlights"`
				} `json:"products"`
			} `json:"mosaic"`
		} `json:"popular"`
		Games struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string    `json:"section_url"`
				SectionType string    `json:"section_type"`
				Layouts     []string  `json:"layouts"`
				Products    []Product `json:"products"`
			} `json:"mosaic"`
		} `json:"games"`
	} `json:"data"`
	UserMaxRewardAmount int    `json:"user_max_reward_amount"`
	CsrfFormKey         string `json:"csrfFormKey"`
}

// SoftwareData has been autogenerated based on the JSON content using https://transform.tools/json-to-go.
type SoftwareData struct {
	UserOptions struct {
		IsLoggedIn bool   `json:"is_logged_in"`
		LogoutURL  string `json:"logout_url"`
	} `json:"userOptions"`
	ShowSingleSignOn           bool `json:"showSingleSignOn"`
	Debug                      bool `json:"debug"`
	IPInChina                  bool `json:"ipInChina"`
	IsEuCountry                bool `json:"isEuCountry"`
	BaseSubscriptionPriceMoney struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"baseSubscriptionPrice|money"`
	BannerOptions interface{} `json:"bannerOptions"`
	ExchangeRates struct {
		USDDecimal float64 `json:"USD|decimal"`
		AUDDecimal float64 `json:"AUD|decimal"`
		CHFDecimal float64 `json:"CHF|decimal"`
		IDRDecimal float64 `json:"IDR|decimal"`
		KRWDecimal float64 `json:"KRW|decimal"`
		BGNDecimal float64 `json:"BGN|decimal"`
		CNYDecimal float64 `json:"CNY|decimal"`
		ISKDecimal float64 `json:"ISK|decimal"`
		ILSDecimal float64 `json:"ILS|decimal"`
		GBPDecimal float64 `json:"GBP|decimal"`
		NZDDecimal float64 `json:"NZD|decimal"`
		DKKDecimal float64 `json:"DKK|decimal"`
		CADDecimal float64 `json:"CAD|decimal"`
		TRYDecimal float64 `json:"TRY|decimal"`
		HUFDecimal float64 `json:"HUF|decimal"`
		PHPDecimal float64 `json:"PHP|decimal"`
		RONDecimal float64 `json:"RON|decimal"`
		NOKDecimal float64 `json:"NOK|decimal"`
		RUBDecimal float64 `json:"RUB|decimal"`
		ZARDecimal float64 `json:"ZAR|decimal"`
		MYRDecimal float64 `json:"MYR|decimal"`
		INRDecimal float64 `json:"INR|decimal"`
		THBDecimal float64 `json:"THB|decimal"`
		MXNDecimal float64 `json:"MXN|decimal"`
		CZKDecimal float64 `json:"CZK|decimal"`
		BRLDecimal float64 `json:"BRL|decimal"`
		JPYDecimal float64 `json:"JPY|decimal"`
		PLNDecimal float64 `json:"PLN|decimal"`
		EURDecimal float64 `json:"EUR|decimal"`
		SEKDecimal float64 `json:"SEK|decimal"`
		SGDDecimal float64 `json:"SGD|decimal"`
		HKDDecimal float64 `json:"HKD|decimal"`
	} `json:"exchangeRates"`
	CsrfTokenInput string `json:"csrfTokenInput"`
	CsrfToken      struct {
	} `json:"csrfToken"`
	Data struct {
		Popular struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string   `json:"section_url"`
				SectionType string   `json:"section_type"`
				Layouts     []string `json:"layouts"`
				Products    []struct {
					TileLogoInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_logo_information"`
					MachineName      string `json:"machine_name"`
					HighResTileImage string `json:"high_res_tile_image"`
					DisableHeroTile  bool   `json:"disable_hero_tile"`
					MarketingBlurb   string `json:"marketing_blurb"`
					HoverTitle       string `json:"hover_title"`
					ProductURL       string `json:"product_url"`
					TileImage        string `json:"tile_image"`
					Category         string `json:"category"`
					HeroHighlights   []struct {
						Heading string `json:"heading"`
						Tooltip string `json:"tooltip"`
					} `json:"hero_highlights"`
					HoverHighlights             []string `json:"hover_highlights"`
					Author                      string   `json:"author"`
					FallbackStoreSaleLogo       string   `json:"fallback_store_sale_logo"`
					HighResTileImageInformation struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"high_res_tile_image_information"`
					SupportsPartners       bool    `json:"supports_partners"`
					DetailedMarketingBlurb string  `json:"detailed_marketing_blurb"`
					TileLogo               string  `json:"tile_logo"`
					TileShortName          string  `json:"tile_short_name"`
					StartDateDatetime      string  `json:"start_date|datetime"`
					EndDateDatetime        string  `json:"end_date|datetime"`
					TileStamp              string  `json:"tile_stamp"`
					BundlesSoldDecimal     float64 `json:"bundles_sold|decimal"`
					TileName               string  `json:"tile_name"`
					ShortMarketingBlurb    string  `json:"short_marketing_blurb"`
					TileImageInformation   struct {
						Config struct {
							ImageType string `json:"image_type"`
							Gcs       string `json:"gcs"`
							Imgix     struct {
								Args struct {
								} `json:"args"`
								MasterImage struct {
									ImageType string `json:"image_type"`
									Gcs       string `json:"gcs"`
									Static    string `json:"static"`
									Imgix     struct {
									} `json:"imgix"`
								} `json:"master_image"`
							} `json:"imgix"`
						} `json:"config"`
					} `json:"tile_image_information"`
					Type       string   `json:"type"`
					Highlights []string `json:"highlights"`
				} `json:"products"`
			} `json:"mosaic"`
		} `json:"popular"`
		Software struct {
			Mosaic []struct {
				SectionURLText string `json:"section_url_text"`
				SectionTitle   struct {
				} `json:"section_title"`
				SectionURL  string    `json:"section_url"`
				SectionType string    `json:"section_type"`
				Layouts     []string  `json:"layouts"`
				Products    []Product `json:"products"`
			} `json:"mosaic"`
		} `json:"software"`
	} `json:"data"`
	UserMaxRewardAmount int    `json:"user_max_reward_amount"`
	CsrfFormKey         string `json:"csrfFormKey"`
}

type Product struct {
	TileLogoInformation struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"tile_logo_information"`
	MachineName      string `json:"machine_name"`
	HighResTileImage string `json:"high_res_tile_image"`
	DisableHeroTile  bool   `json:"disable_hero_tile"`
	MarketingBlurb   string `json:"marketing_blurb"`
	HoverTitle       string `json:"hover_title"`
	ProductURL       string `json:"product_url"`
	TileImage        string `json:"tile_image"`
	Category         string `json:"category"`
	HeroHighlights   []struct {
		Heading string `json:"heading"`
		Tooltip string `json:"tooltip"`
	} `json:"hero_highlights"`
	HoverHighlights             []string `json:"hover_highlights"`
	Author                      string   `json:"author"`
	FallbackStoreSaleLogo       string   `json:"fallback_store_sale_logo"`
	HighResTileImageInformation struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"high_res_tile_image_information"`
	SupportsPartners       bool    `json:"supports_partners"`
	DetailedMarketingBlurb string  `json:"detailed_marketing_blurb"`
	TileLogo               string  `json:"tile_logo"`
	TileShortName          string  `json:"tile_short_name"`
	StartDateDatetime      string  `json:"start_date|datetime"`
	EndDateDatetime        string  `json:"end_date|datetime"`
	TileStamp              string  `json:"tile_stamp"`
	BundlesSoldDecimal     float64 `json:"bundles_sold|decimal"`
	TileName               string  `json:"tile_name"`
	ShortMarketingBlurb    string  `json:"short_marketing_blurb"`
	TileImageInformation   struct {
		Config struct {
			ImageType string `json:"image_type"`
			Gcs       string `json:"gcs"`
			Imgix     struct {
				Args struct {
				} `json:"args"`
				MasterImage struct {
					ImageType string `json:"image_type"`
					Gcs       string `json:"gcs"`
					Static    string `json:"static"`
					Imgix     struct {
					} `json:"imgix"`
				} `json:"master_image"`
			} `json:"imgix"`
		} `json:"config"`
	} `json:"tile_image_information"`
	Type       string   `json:"type"`
	Highlights []string `json:"highlights"`
}
