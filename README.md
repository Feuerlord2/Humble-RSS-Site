# Fanatical-RSS-Site

RSS feeds for Fanatical bundle deals. Queries the Fanatical Algolia API every 6 hours and generates three RSS feeds — books, games, software.

**Live:** https://feuerlord2.github.io/Fanatical-RSS-Site/

## Feeds

```
https://feuerlord2.github.io/Fanatical-RSS-Site/books.rss
https://feuerlord2.github.io/Fanatical-RSS-Site/games.rss
https://feuerlord2.github.io/Fanatical-RSS-Site/software.rss
```

Add these to any RSS reader, Discord bot, or news aggregator. Each item includes current price, original price, discount percentage, and a direct link to the deal.

## How it works

A Go program hits Fanatical's public Algolia API endpoint, converts the response into internal bundle structs, categorizes them (books/games/software based on `display_type` and title matching), deduplicates, and writes RSS 2.0 XML files. GitHub Actions runs this on a schedule and deploys the updated feeds to GitHub Pages.

## Running locally

```
go build -o gofanatical ./cmd/
./gofanatical
```

Requires Go 1.20+. Output goes to `docs/`.

## Project structure

```
cmd/gofanatical.go  Entry point
pkg/feed.go         API fetching, conversion, categorization, RSS generation
pkg/model.go        Data types (FanaticalBundle, Price, Item)
docs/               GitHub Pages output (HTML + RSS files)
Makefile            Build, run, test, serve, docker targets
```

## Disclaimer

Not affiliated with Fanatical. This is a private hobby project.

## License

MIT — see [LICENSE](LICENSE).
