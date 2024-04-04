package reader_test

import (
	"bitcly/click"
	"bitcly/reader"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestPath(t *testing.T) {
	var cases = []struct {
		name     string
		base     string
		fileName string
		expected string
	}{
		{"success", "reader/data", "encoder.json", "reader/data/encoder.json"},
		{"empty", "reader/data", "", "reader/data/"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			path := reader.Path(tt.base, tt.fileName)
			require.Equal(t, tt.expected, path, "Result path is not equal")
		})
	}
}

func TestGetClicks(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2020-02-15T00:00:00Z")
	require.NoError(t, err)
	t2, err := time.Parse(time.RFC3339, "2021-02-15T00:00:00Z")
	require.NoError(t, err)
	t3, err := time.Parse(time.RFC3339, "2022-06-15T00:00:00Z")
	require.NoError(t, err)
	t4, err := time.Parse(time.RFC3339, "2021-10-15T00:00:00Z")
	require.NoError(t, err)

	var cases = []struct {
		name        string
		base        string
		fileName    string
		expectedErr bool
		expected    click.Clicks
	}{
		{"not-found", "testdata", "not-found.json", true, click.Clicks{}},
		{"invalid-JSON", "testdata", "invalid.json", true, click.Clicks{}},
		{
			"valid-JSON",
			"testdata",
			"valid.json",
			false,
			click.Clicks{
				{Bitlink: "http://bit.ly/2kkAHNs", Timestamp: t1},
				{Bitlink: "http://es.pn/3MgVNnZ", Timestamp: t2},
				{Bitlink: "http://bit.ly/2kJdsg8", Timestamp: t3},
				{Bitlink: "http://amzn.to/3C5IIJm", Timestamp: t4},
			}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := reader.GetClicks(tt.base, tt.fileName)
			require.Equal(t, tt.expectedErr, err != nil, "Error expectations not met")
			if tt.expectedErr {
				return
			}
			require.Equal(t, len(tt.expected), len(result), "Result length not equal")
			for index, r := range result {
				require.Equal(t, tt.expected[index].Bitlink, r.Bitlink, "BitLink not equal")
				require.Equal(t, tt.expected[index].Timestamp, r.Timestamp, "Timestamp not equal")
			}
		})
	}
}

func TestGetHashes(t *testing.T) {
	var cases = []struct {
		name        string
		base        string
		fileName    string
		expectedErr bool
		expected    map[string]string
	}{
		{"not-found", "testdata", "not-found.csv", true, map[string]string{}},
		{"invalid", "testdata", "invalid.csv", true, map[string]string{}},
		{"not-enough-cols", "testdata", "not-enough-cols.csv", true, map[string]string{}},
		{"not-enough-rows", "testdata", "not-enough-rows.csv", true, map[string]string{}},
		{"invalid-url", "testdata", "invalid-url.csv", false, map[string]string{"bit.ly/31Tt55y": "https://google.com"}},
		{
			"with-data",
			"testdata",
			"valid.csv",
			false,
			map[string]string{
				"bit.ly/31Tt55y": "https://google.com",
				"bit.ly/2kJO0qS": "https://github.com",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := reader.GetHashes(tt.base, tt.fileName)
			require.Equal(t, tt.expectedErr, err != nil, "Error expectations not met")
			if tt.expectedErr {
				return
			}
			require.Equal(t, len(tt.expected), len(result), "Result length not equal")
			for hash, fullURL := range result {
				val, ok := tt.expected[hash]
				require.True(t, ok, "expected hash key does not exist")
				require.Equal(t, val, fullURL, "Full URLs dont match")
			}
		})
	}
}
