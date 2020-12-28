package main

import "fmt"

func main() {
	input := []int{0, 1, 1, 1, 2, 2, 3, 3, 4, 4}
	fmt.Println(removeDuplicates(input))
	fmt.Println(input)
}

func removeDuplicates(nums []int) int {
	var tmp, count int
	for i, num := range nums {
		if i == 0 {
			tmp = num
			count++
			continue
		}
		if tmp != num {
			nums[count] = num
			tmp = num
			count++
		}
	}
	return count
}
