package main

import "testing"

var testCases = []struct {
	Input  string
	Output int
}{
	{"Hello World", 5},
	{"", 0},
	{"a ", 1},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, testCase := range testCases {
		result := lengthOfLastWord(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
