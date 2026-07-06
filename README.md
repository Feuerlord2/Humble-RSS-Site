# Humble Bundle RSS Feeds

[![Update RSS](https://github.com/Feuerlord2/Humble-RSS-Site/actions/workflows/update_rss.yml/badge.svg)](https://github.com/Feuerlord2/Humble-RSS-Site/actions/workflows/update_rss.yml)

Unofficial RSS feeds for [Humble Bundle](https://www.humblebundle.com/) bundles,
regenerated automatically every few hours by GitHub Actions and served via
GitHub Pages: **<https://feuerlord2.github.io/Humble-RSS-Site/>**

## Feeds

| Category | Feed URL |
| -------- | -------- |
| Books    | <https://feuerlord2.github.io/Humble-RSS-Site/books.rss> |
| Games    | <https://feuerlord2.github.io/Humble-RSS-Site/games.rss> |
| Software | <https://feuerlord2.github.io/Humble-RSS-Site/software.rss> |

## How it works

A small Go program scrapes the JSON embedded in the Humble Bundle category
pages (`books`, `games`, `software`), converts the current bundles into RSS
2.0 feeds, and writes them to `docs/`. The [Update RSS workflow](.github/workflows/update_rss.yml)
runs it on a schedule and commits the result whenever the bundles change.

## Development

```sh
go test ./...                 # run the tests
go run ./cmd/ -out docs       # regenerate the feeds locally
```

## License

[Apache 2.0](LICENSE)
