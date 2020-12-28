package main

import (
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
	{[]int{4, 3, 2, 1, 4}, 16},
	{[]int{1, 2, 1}, 2},
	{[]int{0, 0}, 0},
	{[]int{1, 0}, 0},
}

func TestMaxArea(t *testing.T) {
	for _, testCase := range testCases {
		result := maxArea(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
