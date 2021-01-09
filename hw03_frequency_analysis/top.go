package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

const (
	topLen = 10
)

func Top10(value string) []string {
	var words = make([]string, 0, topLen)
	if len(value) == 0 {
		return words
	}

	dict := getWordsByFrequency(value)
	for k := range dict {
		words = append(words, k)
	}

	sort.Slice(words, func(i, j int) bool {
		return dict[words[i]] > dict[words[j]]
	})

	if len(words) > topLen {
		return words[:topLen]
	}

	return words
}

func getWordsByFrequency(value string) map[string]int {
	words := strings.Fields(value)

	var dict = make(map[string]int)
	for _, s := range words {
		dict[s]++
	}

	return dict
}
