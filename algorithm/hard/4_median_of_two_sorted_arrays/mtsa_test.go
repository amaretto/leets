package main

import "testing"

var testCases = []struct {
	Nums1  []int
	Nums2  []int
	Result float64
}{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{1}, []int{}, 1.0},
	{[]int{}, []int{2}, 2.0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, testCase := range testCases {
		result := findMedianSortedArrays(testCase.Nums1, testCase.Nums2)
		if result != testCase.Result {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
