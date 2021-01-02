package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	var l, r, sum int

	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// add one loop
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l = j + 1
			r = len(nums) - 1
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for nums[l] == nums[l-1] && l < r {
						l++
					}
				}
			}
		}
	}
	return result
}
