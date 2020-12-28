package main

import "testing"

var testCases = []struct {
	Input  int
	Output int
}{
	{2, 2},
	{3, 3},
	{4, 5},
	{35, 14930352},
	{44, 1134903170},
	{45, 1836311903},
}

func TestClimbStairs(t *testing.T) {
	for _, testCase := range testCases {
		result := climbStairs(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
