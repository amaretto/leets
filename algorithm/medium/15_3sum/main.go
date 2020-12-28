package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	sort.Sort(sort.IntSlice(nums))
	var first, second, third, tmp1, tmp2, tmp3, sum int
	for i := 0; i < len(nums)-2; i++ {
		first = nums[i]
		if first > 0 {
			break
		}
		if i != 0 && first == tmp1 {
			tmp1 = nums[i]
			continue
		}
		tmp1 = nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			second = nums[j]
			if first+second > 0 {
				break
			}
			if j != i+1 && second == tmp2 {
				tmp2 = nums[j]
				continue
			}
			tmp2 = nums[j]
			for k := j + 1; k < len(nums); k++ {
				third = nums[k]
				sum = first + second + third
				if k != j+1 && third == tmp3 {
					tmp3 = nums[k]
					continue
				}
				tmp3 = nums[k]

				switch {
				case sum > 0:
					break
				case sum < 0:
					continue
				case sum == 0:
					result = append(result, []int{first, second, third})
					break
				}
			}
		}
	}

	return result
}
