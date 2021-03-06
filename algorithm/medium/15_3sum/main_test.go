package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output [][]int
}{
	// can't compare zero values
	{[]int{}, [][]int(nil)},
	{[]int{0}, [][]int(nil)},
	{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	{[]int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
	{[]int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
	{[]int{1, 2, -2, -1}, [][]int(nil)},
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
}

func TestThreeSum(t *testing.T) {
	for _, testCase := range testCases {
		result := threeSum(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}

func TestThreeSumImprove(t *testing.T) {
	for _, testCase := range testCases {
		result := threeSumImprove(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
