package main

import "testing"

var testCases = []struct {
	Input  int
	Output int
}{
	{4, 2},
	{8, 2},
}

func TestMySqrt(t *testing.T) {
	for _, testCase := range testCases {
		result := mySqrt(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
