package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	numRows int
	want    [][]int
}{
	{1, [][]int{[]int{1}}},
	{5, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}},
}

func TestIsPlindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := generate(testCase.numRows)
		if reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
