package main

import (
	"bitcly/reader"
	"fmt"
	"log/slog"
	"os"
)

const (
	encodesFile = "encodes.csv"
	decodesFile = "decodes.json"
)

func main() {
	hashes, err := reader.GetHashes(encodesFile)
	if err != nil {
		slog.Error("Error reading CSV file:", err)
		os.Exit(-1)
	}

	clicks, err := reader.GetClicks(decodesFile)
	if err != nil {
		slog.Error("Error reading JSON file:", err)
		os.Exit(-1)
	}

	counts, err := clicks.Process(hashes)
	if err != nil {
		slog.Error("Error processing files:", err)
		os.Exit(-1)
	}

	fmt.Println(counts)
}
