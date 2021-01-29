package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	want   []int
}{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{}, 0, []int{-1, -1}},
}

func TestSwapPairs(t *testing.T) {
	for _, testCase := range testCases {
		result := searchRange(testCase.nums, testCase.target)
		if reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
