package main

import "testing"

var testCases = []struct {
	Input   string
	NumRows int
	Output  string
}{
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"PAYPALISHIRING", 1, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 14, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 100, "PAYPALISHIRING"},
}

func TestConvert(t *testing.T) {
	for _, testCase := range testCases {
		result := convert(testCase.Input, testCase.NumRows)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
