package main

import "testing"

var testCases = []struct {
	X      int
	Result bool
}{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPlindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := isPalindrome(testCase.X)
		if result != testCase.Result {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
