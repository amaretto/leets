package main

import "testing"

var testCases = []struct {
	Input  string
	Output int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRoman(t *testing.T) {
	for _, testCase := range testCases {
		result := romanToInt(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
