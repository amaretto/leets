package main

import "testing"

var testCases = []struct {
	Input  int
	Output string
}{
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
	{5, "111221"},
	{6, "312211"},
}

func TestCountAndSay(t *testing.T) {
	for _, testCase := range testCases {
		result := countAndSay(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
