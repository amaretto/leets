package main

import (
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	want   int
}{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 0, -1},
}

func TestSearch(t *testing.T) {
	for _, testCase := range testCases {
		result := search(testCase.nums, testCase.target)
		if result != testCase.want {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}

func TestSearchImprove(t *testing.T) {
	for _, testCase := range testCases {
		result := searchImprove(testCase.nums, testCase.target)
		if result != testCase.want {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
