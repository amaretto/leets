package main

import "testing"

var testCases = []struct {
	Input  []string
	Output string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"doog", "doodle", "doom", "do"}, "do"},
	{[]string{"", "b"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, testCase := range testCases {
		result := longestCommonPrefix(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
