package main

func search(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

func searchImprove(nums []int, target int) int {
	if nums[0] <= target {
		for i := 0; i <= target && i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		tmp := nums[len(nums)-1]
		for i := len(nums) - 1; 0 <= i && target <= nums[i]; i-- {
			if nums[i] > tmp {
				break
			}
			if nums[i] == target {
				return i
			}
			tmp = nums[i]
		}
	}

	return -1
}
