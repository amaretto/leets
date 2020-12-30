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
	//{[]int{}, [][]int{}},
	//{[]int{0}, [][]int{}},
	{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	{[]int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
	{[]int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
	{[]int{1, 2, -2, -1}, [][]int{}},
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
}

//func TestThreeSum(t *testing.T) {
//	for _, testCase := range testCases {
//		result := threeSum(testCase.Input)
//		if !reflect.DeepEqual(result, testCase.Output) {
//			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
//		}
//	}
//}

func TestThreeSum2(t *testing.T) {
	for _, testCase := range testCases {
		result := threeSum2(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
