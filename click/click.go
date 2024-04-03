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

type Click struct {
	Bitlink   string    `json:"bitlink"`
	Timestamp time.Time `json:"timestamp"`
}

type Clicks []Click

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
