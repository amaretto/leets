package main

import (
	"testing"
)

var testCases = []struct {
	dividend int
	divisor  int
	want     int
}{
	{10, 3, 3},
	{7, -3, -2},
	{-64, 5, -12},
	{-64, -12, 5},
	{-2147483648, -1, 2147483647},
}

func TestSwapPairs(t *testing.T) {
	for _, testCase := range testCases {
		result := divide(testCase.dividend, testCase.divisor)
		if result != testCase.want {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
