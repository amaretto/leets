package main

import "testing"

var testCases = []struct {
	X      int
	Result int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
}

func TestReverse(t *testing.T) {
	for _, testCase := range testCases {
		result := Reverse(testCase.X)
		if result != testCase.Result {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
