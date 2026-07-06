package main

import (
	"flag"
	"log/slog"
	"os"

	gohumble "github.com/Feuerlord2/Humble-RSS-Site/pkg"
)

func main() {
	outDir := flag.String("out", ".", "directory to write the .rss files into")
	flag.Parse()

	if err := gohumble.Run(*outDir); err != nil {
		slog.Error("updating feeds failed", "error", err)
		os.Exit(1)
	}
}
