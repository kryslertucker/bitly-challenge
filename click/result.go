package click

import "sort"

type Result struct {
	LongURL string
	Count   int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type Results []Result

func (r Results) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Results) Len() int           { return len(r) }
func (r Results) Less(i, j int) bool { return r[i].Count < r[j].Count }

func Prepare(countsPerURL map[string]int) Results {
	counts := Results{}
	for longUrl, count := range countsPerURL {
		counts = append(counts, Result{longUrl, count})
	}

	sort.Sort(sort.Reverse(counts))
	return counts
}
