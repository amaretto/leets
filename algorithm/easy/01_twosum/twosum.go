package main

// TwoSum is
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumHash is
func TwoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if m[d] == 0 || m[d] == i {
			continue
		}
		return []int{i, m[d]}
	}
	return nil
}

// TwoSumOnePass is
func TwoSumOnePass(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if _, ok := m[d]; ok {
			return []int{m[d], i}
		}
		m[nums[i]] = i
	}
	return nil
}
