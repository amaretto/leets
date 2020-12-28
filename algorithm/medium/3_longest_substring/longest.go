package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	var count, max int
	start := 1
	m := make(map[rune]int)
	for i, c := range s {
		count++
		if m[c] != 0 {
			if m[c] > start {
				start = m[c]
			}
			count = i + 1 - start
		}
		if count > max {
			max = count
		}
		m[c] = i + 1
	}
	return max
}
