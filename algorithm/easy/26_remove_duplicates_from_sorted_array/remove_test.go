package main

import "testing"

var testCases = []struct {
	Input  []int
	Output int
}{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, testCase := range testCases {
		result := removeDuplicates(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
