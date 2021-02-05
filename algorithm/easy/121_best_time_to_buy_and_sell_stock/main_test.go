package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	prices []int
	want   int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestIsPlindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := maxProfit(testCase.prices)
		if !reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
