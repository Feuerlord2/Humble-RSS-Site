# Humble Watch

Humble Watch is a personal Humble Bundle feed watcher and static dashboard.

It is based on [`Feuerlord2/Humble-RSS-Site`](https://github.com/Feuerlord2/Humble-RSS-Site), originally created by Daniel Winter / Feuerlord2, and extends the original RSS generator with additional JSON outputs, bundle end dates, time-left calculations, and a table-based frontend.

This project is not affiliated with Humble Bundle, Inc.

## What it does

The generator fetches the public Humble Bundle category pages:

- `https://www.humblebundle.com/books`
- `https://www.humblebundle.com/games`
- `https://www.humblebundle.com/software`

It extracts the embedded `script#landingPage-json-data` payload, parses the bundle product data, and generates RSS and JSON outputs in `docs/`.

## Generated outputs

The current generated files are:

- `docs/books.rss` — book bundle RSS feed
- `docs/games.rss` — game bundle RSS feed
- `docs/software.rss` — software bundle RSS feed
- `docs/all.rss` — combined RSS feed
- `docs/all.json` — normalized machine-readable list of all detected bundles
- `docs/urgent.json` — bundles expiring within short time windows

The JSON output includes fields such as:

- `category`
- `title`
- `url`
- `description`
- `start_at`
- `end_at`
- `days_left`
- `hours_left`
- `machine_name`
- `bundles_sold`
- `tile_stamp`

## Static dashboard

`docs/index.html` provides a static table frontend over `all.json`.

It includes:

- category filters,
- urgent `<72h` filter,
- search,
- sortable table columns,
- summary counters,
- links to generated RSS and JSON files.

When GitHub Pages is configured for `main` + `/docs`, the dashboard can be served as a static site.

## Workflows

The repository includes GitHub Actions workflows for:

- Go test/build validation,
- manual feed generation,
- scheduled feed generation around 20:10 Europe/Warsaw.

The scheduled workflow uses UTC cron entries plus a Warsaw-local time gate to handle CEST/CET changes.

## Local usage

Requirements:

- Go 1.20+

Build and generate feeds locally:

```bash
go build -o gohumble ./cmd/
cd docs
../gohumble
```

Then open:

```text
docs/index.html
```

or inspect generated files:

```text
docs/all.json
docs/urgent.json
docs/all.rss
```

## Attribution

This project is a derivative/fork of:

- Original project: [`Feuerlord2/Humble-RSS-Site`](https://github.com/Feuerlord2/Humble-RSS-Site)
- Original author: Daniel Winter / Feuerlord2
- Original inspiration noted upstream: `shimst3r/go-humble`

Additional changes in this fork are maintained under `ThresholdOps/humble-watch`.

## License

This project retains the upstream Apache License 2.0 license. See [`LICENSE`](./LICENSE) and [`NOTICE`](./NOTICE).
