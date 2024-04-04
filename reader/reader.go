// Package reader provides the functionality to read and parse CSV and JSON files
package reader

import (
	"bitcly/click"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"
)

var ErrInvalidFile = errors.New("file is not valid")

// Path returns the full path of the files given a base path and a file name
func Path(base, name string) string {
	return fmt.Sprintf("%s/%s", base, name)
}

// GetHashes reads a CSV file containing the mapping between full
// url values, its domain and associated hash. It returns
// a map where the value is the full URL and the key is the combination
// of the domain and hash.
func GetHashes(base, fname string) (map[string]string, error) {
	file, err := os.Open(Path(base, fname))
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", fname, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV file '%s': %w", fname, err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file must have at least 1 row of values: %w", ErrInvalidFile)
	}

	hashes := make(map[string]string)
	for _, record := range rows[1:] {
		if len(record) < 3 {
			return nil, fmt.Errorf("row must have at least three columns: %w", ErrInvalidFile)
		}
		fullHash := fmt.Sprintf("%s/%s", record[1], record[2])

		_, ok := hashes[fullHash]
		if ok {
			slog.Debug("Duplicate domain and hash combination, skipping record", "value", fullHash)
			continue
		}
		_, err := url.ParseRequestURI(record[0])
		if err != nil {
			slog.Error("Invalid URL, it won't be processed", "url", record[0])
			continue
		}

		hashes[fullHash] = strings.Trim(record[0], "/")
	}

	return hashes, nil
}

// GetClicks reads a JSON file with URL click events and the timestamp of each event
func GetClicks(base, fname string) (click.Clicks, error) {
	raw, err := os.ReadFile(Path(base, fname))
	if err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", fname, err)
	}

	clicks := click.Clicks{}
	if err := json.Unmarshal(raw, &clicks); err != nil {
		return nil, fmt.Errorf("file content can not be parsed: %w", err)
	}

	return clicks, nil
}
