package click_test

import (
	"bitcly/click"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestProcess(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2020-02-15T00:00:00Z")
	require.NoError(t, err)
	t2, err := time.Parse(time.RFC3339, "2021-02-15T00:00:00Z")
	require.NoError(t, err)
	t3, err := time.Parse(time.RFC3339, "2022-06-15T00:00:00Z")
	require.NoError(t, err)
	t4, err := time.Parse(time.RFC3339, "2021-10-15T00:00:00Z")
	require.NoError(t, err)
	t5, err := time.Parse(time.RFC3339, "2021-03-15T00:00:00Z")
	require.NoError(t, err)

	var cases = []struct {
		name        string
		clicks      click.Clicks
		hashes      map[string]string
		expectedErr bool
		expected    click.Results
	}{
		{"no-data", click.Clicks{}, map[string]string{}, false, click.Results{}},
		{
			"invalid-url",
			click.Clicks{
				{Bitlink: "bit.ly/2kkAHNs", Timestamp: t2},
			},
			map[string]string{
				"bit.ly/2kkAHNs": "https://twitter.com",
			},
			false,
			click.Results{},
		},
		{
			"with-data",
			click.Clicks{
				{Bitlink: "http://bit.ly/2kkAHNs", Timestamp: t1},
				{Bitlink: "http://es.pn/3MgVNnZ", Timestamp: t2},
				{Bitlink: "http://bit.ly/2kkAHNs", Timestamp: t2},
				{Bitlink: "http://bit.ly/2kJdsg8", Timestamp: t2},
				{Bitlink: "http://bit.ly/2kJdsg8", Timestamp: t3},
				{Bitlink: "http://amzn.to/3C5IIJm", Timestamp: t4},
				{Bitlink: "http://bit.ly/2kkAHNs", Timestamp: t5},
			},
			map[string]string{
				"bit.ly/2kkAHNs": "https://twitter.com",
				"bit.ly/2kJdsg8": "https://reddit.com",
			},
			false,
			click.Results{
				{"https://twitter.com", 2},
				{"https://reddit.com", 1},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			results, err := tt.clicks.Process(tt.hashes)
			require.Equal(t, tt.expectedErr, err != nil, "Error expectations not met")
			if tt.expectedErr {
				return
			}
			require.Equal(t, len(tt.expected), len(results), "Results length are not equal")
			require.Equal(t, tt.expected, results, "Results are not equal")
		})
	}
}
