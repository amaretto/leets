package main

import (
	"testing"
)

var testCases = []struct {
	Input  int
	Output string
}{
	{1, "I"},
	{3, "III"},
	{4, "IV"},
	{9, "IX"},
	{58, "LVIII"},
	{444, "CDXLIV"},
	{1994, "MCMXCIV"},
	{3999, "MMMCMXCIX"},
}

func TestMaxArea(t *testing.T) {
	for _, testCase := range testCases {
		result := intToRoman(testCase.Input)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
