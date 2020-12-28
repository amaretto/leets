package main

import "testing"

var testCases = []struct {
	Nums   []int
	Val    int
	Output int
}{
	{[]int{3, 2, 2, 3}, 3, 2},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
}

func TestHelloWorld(t *testing.T) {
	for _, testCase := range testCases {
		result := removeElement(testCase.Nums, testCase.Val)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
