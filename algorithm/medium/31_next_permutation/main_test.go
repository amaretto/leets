package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input []int
	want  []int
}{
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 1, 5}, []int{1, 5, 1}},
	{[]int{1}, []int{1}},
}

func TestSwapPairs(t *testing.T) {
	for _, testCase := range testCases {
		result := testCase.input
		nextPermutation(result)
		if !reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
