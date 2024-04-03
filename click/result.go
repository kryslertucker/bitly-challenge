package click

import (
	"fmt"
	"sort"
	"strings"
)

type Result struct {
	LongURL string
	Count   int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type Results []Result

func (r Results) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Results) Len() int           { return len(r) }
func (r Results) Less(i, j int) bool { return r[i].Count < r[j].Count }

func (r Results) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for i, res := range r {
		builder.WriteString(res.String())
		if i == len(r)-1 {
			break
		}
		builder.WriteString(", ")
	}
	builder.WriteString("]")
	return builder.String()
}

func (res Result) String() string {
	return fmt.Sprintf("{\"%s\": %d}", res.LongURL, res.Count)
}

func Prepare(countsPerURL map[string]int) Results {
	counts := Results{}
	for longUrl, count := range countsPerURL {
		counts = append(counts, Result{longUrl, count})
	}

	sort.Sort(sort.Reverse(counts))
	return counts
}
