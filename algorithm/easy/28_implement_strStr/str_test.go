package main

import "testing"

var testCases = []struct {
	Haystack string
	Needle   string
	Output   int
}{
	{"hello", "ll", 2},
	{"aaaaa", "bba", -1},
	{"aaaaa", "", 0},
	{"a", "a", 0},
	{"mississippi", "issip", 4},
}

func TestHelloWorld(t *testing.T) {
	for _, testCase := range testCases {
		result := strStr(testCase.Haystack, testCase.Needle)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
