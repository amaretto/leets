package main

import "testing"

var testCases = []struct {
	Nums   []int
	Target int
	Output int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, testCase := range testCases {
		result := searchInsert(testCase.Nums, testCase.Target)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
