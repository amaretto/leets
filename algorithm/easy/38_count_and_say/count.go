package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	result := "1"
	tmp := ""
	sc := 0
	var tmpc string

	for i := 1; i < n; i++ {
		tmp = result
		result = ""
		for _, c := range tmp {
			if tmpc == string(c) {
				sc++
			} else {
				if sc == 1 {
					if tmpc == "1" {
						result += "11"
					} else {
						result += "1" + tmpc
					}
				} else if sc > 1 {
					result += strconv.Itoa(sc) + tmpc
				}
				sc = 1
			}
			tmpc = string(c)
		}
		if sc == 1 {
			if tmpc == "1" {
				result += "11"
			} else {
				result += "1" + tmpc
			}
		} else if sc > 1 {
			if tmpc == "1" {
				result += strconv.Itoa(sc) + "1"
			} else {
				result += strconv.Itoa(sc) + tmpc
			}
		}
		sc = 0
		tmpc = ""
	}
	return result
}
