// Package click provides a set of functions and methods to
// manipulate Clicks and Results
package click

import (
	"fmt"
	"strings"
	"time"
)

const (
	startDate  = "2021-01-01T00:00:00Z"
	cutOffDate = "2022-01-01T00:00:00Z"
)

// Click represents a click on a link at a moment on time
type Click struct {
	Bitlink   string    `json:"bitlink"`
	Timestamp time.Time `json:"timestamp"`
}

// Clicks wraps multiple clicks
type Clicks []Click

// Process iterates over each instance of a click. It skips the event if it
// ocurred outside of 2021 year. Otherwise, it verifies the url
// clicked exists as a key in the provided hash passed as argument, and
// if it does increases the count of the longUrl.
//
// It returns a slice of Results sorted by count value in descending order
func (clicks Clicks) Process(hashes map[string]string) (Results, error) {
	counts := make(map[string]int, len(hashes))
	for _, click := range clicks {
		startTime, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			return nil, fmt.Errorf("could not parse cut off date '%s': %w", cutOffDate, err)
		}

		cutOffTime, err := time.Parse(time.RFC3339, cutOffDate)
		if err != nil {
			return nil, fmt.Errorf("could not parse cut off date '%s': %w", cutOffDate, err)
		}
		if click.Timestamp.Before(startTime) || click.Timestamp.After(cutOffTime) {
			// TODO: log here
			continue
		}

		cleanLink := strings.ReplaceAll(click.Bitlink, "http://", "")
		cleanLink = strings.ReplaceAll(cleanLink, "https://", "")

		longUrl, ok := hashes[cleanLink]
		if !ok {
			// TODO: log here
			continue
		}
		counts[longUrl]++
	}

	return Prepare(counts), nil
}
