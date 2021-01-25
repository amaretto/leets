package main

import (
	"reflect"
)

func findSubstring(s string, words []string) []int {
	var result []int
	if s == "" || words == nil {
		return result
	}

	// create map[word]int
	wlength := len(words[0])
	clength := len(words[0]) * len(words)
	mlength := distinct(words)
	checkList := makeCheckList(words)

	for i := 0; i+clength <= len(s); i++ {
		tmp := s[i : i+clength]
		countMap := make(map[string]int, mlength)
		for j := 0; j+wlength <= len(tmp); j += wlength {
			word := tmp[j : j+wlength]
			if contains(word, words) {
				countMap[word]++
			} else {
				break
			}
			if len(checkList) == len(countMap) && reflect.DeepEqual(checkList, countMap) {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func makeCheckList(words []string) map[string]int {
	m := make(map[string]int, distinct(words))
	for _, w := range words {
		m[w]++
	}
	return m
}

func contains(word string, words []string) bool {
	for _, w := range words {
		if word == w {
			return true
		}
	}
	return false
}

func distinct(words []string) int {
	m := make(map[string]bool, len(words))
	for _, w := range words {
		m[w] = true
	}
	return len(m)
}
