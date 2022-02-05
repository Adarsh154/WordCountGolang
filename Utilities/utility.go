package Utilities

import (
	"sort"
	"strings"
)

type Pair struct {
	Word  string
	Count int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl[:10]
}

func Repetition(st string) PairList {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		if _, ok := wc[word]; ok {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return rankByWordCount(wc)
}
