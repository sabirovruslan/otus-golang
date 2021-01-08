package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

const topLen = 10

func Top10(value string) []string {
	var top = make([]string, 0, topLen)
	if len(value) == 0 {
		return top
	}

	dict := getWordsByFrequency(value)
	for k := range dict {
		top = append(top, k)
	}

	sort.Slice(top, func(i, j int) bool {
		return dict[top[i]] > dict[top[j]]
	})

	if len(top) > topLen {
		return top[:topLen]
	}

	return top
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
