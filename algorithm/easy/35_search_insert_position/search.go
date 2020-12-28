package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}
