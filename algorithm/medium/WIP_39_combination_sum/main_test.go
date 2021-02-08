package main

import (
	"fmt"
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

func TestFunction(t *testing.T) {
	for _, testCase := range testCases {
		result := combinationSum(testCase.candidates, testCase.target)
		fmt.Println(result)
		if !isSame(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}

func isSame(a, b [][]int) bool {
	for _, i := range a {
		for _, j := range b {
			reflect.DeepEqual(i, j)
			continue
		}
		return false
	}
	return true
}
