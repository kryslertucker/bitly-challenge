package main

import (
	"bitcly/reader"
	"fmt"
)

const (
	encodesFile = "encodes.csv"
	decodesFile = "decodes.json"
)

func main() {
	hashes, err := reader.GetHashes(encodesFile)
	if err != nil {
		// fmt.Errorf("could not read CSV file '%s': %w", encodesFile, err)
		return
	}

	clicks, err := reader.GetClicks(decodesFile)
	if err != nil {
		// TODO: error here
		return
	}

	counts, err := clicks.Process(hashes)
	if err != nil {
		// TODO: error here
		return
	}

	fmt.Println(counts)

}
