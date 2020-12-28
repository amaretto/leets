package main

import "testing"

var testCases = []struct {
	Input  string
	Output int
}{
	{"42", 42},
	{"   -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
	{"-91283472332", -2147483648},
	{"+1", 1},
	{"+-2", 0},
	{"-0012a42", -12},
}

func TestMyAtoi(t *testing.T) {
	for _, testCase := range testCases {
		result := myAtoi(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
