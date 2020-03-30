package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(in string) []string {
	in = regexp.MustCompile(`\n`).ReplaceAllString(in, " ")
	in = regexp.MustCompile(`\t`).ReplaceAllString(in, "")
	in = regexp.MustCompile(`\s\s+`).ReplaceAllString(in, " ")

	if in == "" {
		return make([]string, 0)
	}

	s1 := strings.Split(in, " ")
	m := make(map[string]int)

	// feed slice into map for counting
	for _, word := range s1 {
		count := m[word]
		m[word] = count + 1 //nolint:gomnd
	}

	// create new slice out of unique words
	s2 := make([]string, 0, len(m))
	for key := range m {
		s2 = append(s2, key)
	}

	// sort based on map values
	sort.Slice(s2, func(i, j int) bool { return m[s2[i]] > m[s2[j]] })

	// return slice up to 10 elements
	top := 10
	if len(m) <= top {
		return s2
	}
	return s2[:top]
}
