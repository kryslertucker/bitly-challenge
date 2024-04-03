package reader

import (
	"bitcly/click"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	base = "reader/data"
)

var ErrInvalidFile = errors.New("file is not valid")

func Path(name string) string {
	return fmt.Sprintf("%s/%s", base, name)
}

func CSV(name string) ([][]string, error) {
	file, err := os.Open(Path(name))
	if err != nil {
		// TODO: error here
		return [][]string{}, fmt.Errorf("unable to open file '%s': %w", name, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, fmt.Errorf("could not read CSV file '%s': %w", name, err)
	}

	if len(rows) < 2 {
		return [][]string{}, fmt.Errorf("file must have at least 1 row of values: %w", ErrInvalidFile)
	}

	return rows, nil
}

func JSON(name string) ([]byte, error) {
	raw, err := os.ReadFile(Path(name))
	if err != nil {
		return []byte{}, fmt.Errorf("error reading JSON file '%s': %w", name, err)
	}

	return raw, nil
}

func GetHashes(fname string) (map[string]string, error) {
	rows, err := CSV(fname)
	if err != nil {
		return nil, err
	}

	hashes := make(map[string]string)
	for _, record := range rows {
		if len(record) < 2 {
			return nil, fmt.Errorf("row must have at least two columns: %w", ErrInvalidFile)
		}
		fullHash := fmt.Sprintf("%s/%s", record[1], record[2])

		_, ok := hashes[fullHash]
		if ok {
			// TODO: log dupe
			continue
		}
		hashes[fullHash] = record[0]
	}

	return hashes, nil
}

func GetClicks(fname string) (click.Clicks, error) {
	decodes, err := JSON(fname)
	if err != nil {
		return nil, err
	}

	clicks := click.Clicks{}
	if err := json.Unmarshal(decodes, &clicks); err != nil {
		return nil, fmt.Errorf("file content can not be parsed: %w", err)
	}

	return clicks, nil
}
