package main

import "testing"

var testCases = []struct {
	Input  string
	Output string
}{
	{"bb", "bb"},
	{"abcba", "abcba"},
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"a", "a"},
	{"", ""},
	{"aaaa", "aaaa"},
	{"abcda", "a"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := longestPalindrome(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
