package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	rowIndex int
	want     []int
}{
	{0, []int{1}},
	{1, []int{1, 1}},
	{3, []int{1, 3, 3, 1}},
	{4, []int{1, 4, 6, 4, 1}},
}

func TestIsPlindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := getRow(testCase.rowIndex)
		if !reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
