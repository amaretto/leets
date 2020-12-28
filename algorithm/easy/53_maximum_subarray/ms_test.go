package main

import "testing"

var testCases = []struct {
	Inputs []int
	Output int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{-1}, -1},
}

func TestMaxSubArray(t *testing.T) {
	for _, testCase := range testCases {
		result := maxSubArray(testCase.Inputs)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
