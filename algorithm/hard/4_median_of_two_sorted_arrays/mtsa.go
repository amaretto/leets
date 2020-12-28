package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tl := len(nums1) + len(nums2)
	limit := tl/2 + 1
	var pos1, pos2 int
	var isOdd bool

	if tl%2 == 1 {
		isOdd = true
	} else {
		isOdd = false
	}

	var tmp, pre int
	for i := 0; i < limit; i++ {
		if pos1 == len(nums1) {
			pre = tmp
			tmp = nums2[pos2]
			pos2++
			continue
		} else if pos2 == len(nums2) {
			pre = tmp
			tmp = nums1[pos1]
			pos1++
			continue
		}

		if nums1[pos1] < nums2[pos2] {
			pre = tmp
			tmp = nums1[pos1]
			pos1++
		} else {
			pre = tmp
			tmp = nums2[pos2]
			pos2++
		}
	}

	if isOdd {
		return float64(tmp)
	}
	return (float64(pre) + float64(tmp)) / 2
}
