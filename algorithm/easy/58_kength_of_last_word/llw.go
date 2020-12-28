package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("a "))
}

// use less memory
func lengthOfLastWord(s string) int {
	var flag bool
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if !flag {
			if s[i] == 32 {
				continue
			}
			flag = true
			count = 1
		} else {
			if s[i] == 32 {
				return count
			}
			count++
		}
	}
	return count
}

// memory usage is too big
func lengthOfLastWord2(s string) int {
	sa := strings.Split(s, " ")

	for i := len(sa) - 1; i >= 0; i-- {
		if sa[i] != "" {
			return len(sa[i])
		}
	}

	return 0
}
