package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	var lidx, count int
	for _, num := range nums {
		if num != val {
			count++
			nums[lidx] = num
			lidx++
		}
	}
	return count
}
