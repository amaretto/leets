package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// accept
func accesptMaxSubArray(nums []int) int {
	var tmp, max int

	max = nums[0]

	for i := 0; i < len(nums); i++ {
		tmp = 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

// timeout
func timeoutMaxSubArray(nums []int) int {
	var sa []int
	var tmp, max int

	max = nums[0]

	for i := 1; i <= len(nums); i++ {
		for j := 0; j+i <= len(nums); j++ {
			sa = nums[j : j+i]
			tmp = 0
			for _, num := range sa {
				tmp += num
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
