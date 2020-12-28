package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	var sum, adv int
	adv = 1
	result := make([]int, len(digits))
	for i := len(digits); 0 < i; i-- {
		sum = digits[i-1] + adv
		if sum == 10 {
			// last adv
			if i == 1 {
				result[i-1] = 0
				result = append([]int{1}, result...)
				break
			}
			result[i-1] = 0
			adv = 1
			continue
		}
		result[i-1] = sum
		adv = 0
	}
	return result
}
