package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

var rx = regexp.MustCompile(`\s\s+`)

func Top10(in string) []string {
	in = strings.ReplaceAll(in, `\n`, " ")
	in = strings.ReplaceAll(in, `\t`, "")
	in = rx.ReplaceAllString(in, " ")

	if in == "" {
		return nil
	}

	s1 := strings.Split(in, " ")
	m := make(map[string]int)

	// feed slice into map for counting
	for _, word := range s1 {
		m[word]++
	}

	// create new slice out of unique words
	s2 := make([]string, 0, len(m))
	for key := range m {
		s2 = append(s2, key)
	}

	// sort based on map values
	sort.Slice(s2, func(i, j int) bool { return m[s2[i]] > m[s2[j]] })

	// return slice up to 10 elements
	const top = 10
	if len(m) <= top {
		return s2
	}
	return s2[:top]
}
