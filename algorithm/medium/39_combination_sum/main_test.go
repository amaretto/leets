package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	candidates []int
	target     int
	want       [][]int
}{
	{[]int{2, 3, 6, 7}, 7, [][]int{[]int{2, 2, 3}, []int{7}}},
	{[]int{2, 3, 5}, 8, [][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}}},
	{[]int{2}, 1, [][]int(nil)},
	{[]int{1}, 1, [][]int{[]int{1}}},
}

func TestCombinationSum(t *testing.T) {
	for _, testCase := range testCases {
		result := combinationSum(testCase.candidates, testCase.target)
		if !isSame(testCase.want, result) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}

func isSame(a, b [][]int) bool {
	var isContain bool
	for _, i := range a {
		isContain = false
		for _, j := range b {
			if reflect.DeepEqual(i, j) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}

	return true
}
