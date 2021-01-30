package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums)-1
	var mid int

	for l < r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
			continue
		} else if target < nums[mid] {
			r = mid - 1
			continue
		} else {
			l, r = mid, mid
			for {
				if l-1 < 0 || nums[l-1] != target {
					break
				}
				l--
			}
			for {
				if r+1 < len(nums) && nums[r+1] != target {
					break
				}
				r++
			}
			break
		}
	}
	if r < l || ((r == l) && nums[l] != target) {
		return []int{-1, -1}
	}
	return []int{l, r}
}
