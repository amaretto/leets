package main

import "testing"

var testCases = []struct {
	Input  string
	Output int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"abba", 2},
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	for _, testCase := range testCases {
		result := lengthOfLongestSubstring(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
