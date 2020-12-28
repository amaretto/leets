package main

import "testing"

var testCases = []struct {
	Input  string
	Output bool
}{
	{"()", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
	{"()[]{}", true},
	{"((", false},
}

func TestIsValid(t *testing.T) {
	for _, testCase := range testCases {
		result := isValid(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
