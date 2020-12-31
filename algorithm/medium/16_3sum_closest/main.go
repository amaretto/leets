package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	var l, r, sum, ans, dist int
	mindis := 2000

	for i := 0; i < len(nums); i++ {
		l, r = i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			dist = distance(sum, target)

			if dist < mindis {
				ans = sum
				mindis = dist
			}

			if sum < target {
				l++
			} else if target < sum {
				r--
			} else {
				return target
			}
		}
	}
	return ans
}

func distance(x, y int) int {
	result := x - y
	if result < 0 {
		return result * -1
	}
	return result
}
