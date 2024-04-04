package click

import (
	"fmt"
	"sort"
	"strings"
)

// Result defines the number of clicks per URL
type Result struct {
	LongURL string
	Count   int
}

// Results wraps a set of click counts per URL and implements sort.Interface to sort by Count.
type Results []Result

func (r Results) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Results) Len() int           { return len(r) }
func (r Results) Less(i, j int) bool { return r[i].Count < r[j].Count }

func (res Result) Empty() bool {
	return res.LongURL == "" && res.Count == 0
}

// String is a custom implementation to be invoked when the of results
// is converted to a string. It returns a string of an array of
// objects separated by comma.
func (r Results) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "[")
	for i, res := range r {
		if res.Empty() {
			continue
		}
		fmt.Fprint(&builder, res.String())
		if i == len(r)-1 {
			break
		}
		fmt.Fprintf(&builder, ", ")
	}
	fmt.Fprintf(&builder, "]")
	return builder.String()
}

// String is the custom implementation of a result where
// the object returned has a key being the longURL
// and value being the count of clicks
func (res Result) String() string {
	if res.Empty() {
		return ""
	}
	return fmt.Sprintf("{\"%s\": %d}", res.LongURL, res.Count)
}

// Prepare iterates over each element in the map
// representing the number of clicks per URL and returns
// a slice sorted by count in descending order
func Prepare(countsPerURL map[string]int) Results {
	counts := Results{}
	for longUrl, count := range countsPerURL {
		counts = append(counts, Result{longUrl, count})
	}

	sort.Sort(sort.Reverse(counts))
	return counts
}
