package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Target int
	Output int
}{
	{[]int{0, 0, 0}, 10, 0},
	{[]int{-1, 2, 1, -4}, 1, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 10},
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 10, 10},
}

func TestThreeSumClosest(t *testing.T) {
	for _, testCase := range testCases {
		result := threeSumClosest(testCase.Input, testCase.Target)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
