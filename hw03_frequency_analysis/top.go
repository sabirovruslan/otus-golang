package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
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
	words := regexp.MustCompile("[ \n\t]").Split(value, -1)

	var dict = make(map[string]int)
	for _, s := range words {
		s = strings.Trim(s, " ")
		if len(s) == 0 {
			continue
		}
		dict[s]++
	}

	return dict
}
