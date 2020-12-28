package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := 0
	for i, num := range nums1 {
		if idx == n {
			break
		}
		if m == 0 && i == 0 {
			nums1[i] = nums2[idx]
			idx++
			continue
		}
		if nums2[idx] < num || m <= i {
			for j := len(nums1) - 1; i < j; j-- {
				nums1[j] = nums1[j-1]
			}
			nums1[i] = nums2[idx]
			m++
			idx++
		}
	}
}
