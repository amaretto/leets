package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Target int
	Output [][]int
}{
	{[]int{}, 0, [][]int(nil)},
	{[]int{0}, 0, [][]int(nil)},
	{[]int{0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
	{[]int{0, 0, 0, 0}, 1, [][]int(nil)},
	{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	{[]int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
	{[]int{-2, -1, 0, 1, 2}, 0, [][]int{{-2, -1, 1, 2}}},
}

func TestThreeSumImprove(t *testing.T) {
	for _, testCase := range testCases {
		result := fourSum(testCase.Input, testCase.Target)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
