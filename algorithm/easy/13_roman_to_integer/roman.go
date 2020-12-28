package main

import "fmt"

func main() {
	var tmp rune
	fmt.Println(tmp == 0)
}

func romanToInt(s string) int {
	var symbol = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int
	var tmp, neu int
	for _, c := range s {
		neu = symbol[c]
		if tmp == 0 {
			tmp = neu
		} else if neu > tmp {
			sum += neu - tmp
			tmp = 0
		} else {
			sum += tmp
			tmp = neu
		}
	}
	if tmp == 0 {
		return sum
	}
	return sum + neu
}
