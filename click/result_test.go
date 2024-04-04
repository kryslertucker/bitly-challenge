package click_test

import (
	"bitcly/click"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResultString(t *testing.T) {
	var cases = []struct {
		name     string
		result   click.Result
		expected string
	}{
		{"no-data", click.Result{}, ""},
		{"data", click.Result{"https://google.com/", 10}, "{\"https://google.com/\": 10}"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.result.String(), "Resul string is not equal")
		})
	}
}

func TestResultsString(t *testing.T) {
	var cases = []struct {
		name     string
		results  click.Results
		expected string
	}{
		{"no-data", click.Results{}, "[]"},
		{
			"data",
			click.Results{
				{"https://google.com/", 10},
				{"https://reddit.com/", 9},
				{"https://linkedin.com/", 2},
			},
			"[{\"https://google.com/\": 10}, {\"https://reddit.com/\": 9}, {\"https://linkedin.com/\": 2}]",
		},
		{
			"data-with-empty",
			click.Results{
				{"https://google.com/", 10},
				{"", 0},
				{"https://reddit.com/", 9},
				{"https://linkedin.com/", 2},
			},
			"[{\"https://google.com/\": 10}, {\"https://reddit.com/\": 9}, {\"https://linkedin.com/\": 2}]",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.results.String(), "Results string is not equal")
		})
	}
}

func TestPrepare(t *testing.T) {
	var cases = []struct {
		name     string
		counts   map[string]int
		expected click.Results
	}{
		{"no-data", map[string]int{}, click.Results{}},
		{"some-data", map[string]int{
			"https://google.com/":  494,
			"https://youtube.com/": 558,
			"https://reddit.com/":  511,
		}, click.Results{
			{"https://youtube.com/", 558},
			{"https://reddit.com/", 511},
			{"https://google.com/", 494},
		}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := click.Prepare(tt.counts)
			require.Equal(t, len(tt.expected), len(got), "Results counts are not equal")
			require.Equal(t, tt.expected, got, "Results values are not equal")
		})
	}
}
