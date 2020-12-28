package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Nums1  []int
	M      int
	Nums2  []int
	N      int
	Output []int
}{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	{[]int{1}, 1, []int{}, 0, []int{1}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
	{[]int{1, 0}, 1, []int{2}, 1, []int{1, 2}},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		merge(testCase.Nums1, testCase.M, testCase.Nums2, testCase.N)
		if !reflect.DeepEqual(testCase.Nums1, testCase.Output) {
			t.Errorf("invalid result. testCase:%v", testCase)
		}
	}
}
