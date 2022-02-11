package Utilities

import (
	"log"
	"regexp"
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

// rankByWordCount finds the top 10 repeated words from the word count map
func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Repetition Calculates word count in a map
func Repetition(st string) PairList {
	// Processing string before counting
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(st, " ")
	SliceString := strings.Fields(processedString)

	//Word count
	wc := make(map[string]int)
	for _, word := range SliceString {
		word = strings.ToLower(word)
		if _, ok := wc[word]; ok {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return rankByWordCount(wc)
}
