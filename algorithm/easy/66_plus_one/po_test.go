package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	{[]int{1, 9, 9}, []int{2, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	for _, testCase := range testCases {
		result := plusOne(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
