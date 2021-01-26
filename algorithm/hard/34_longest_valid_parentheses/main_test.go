package main

import (
	"testing"
)

var testCases = []struct {
	s    string
	want int
}{
	{"(()", 2},
	{")()())", 4},
	{"(())", 4},
	{"(())()()", 8},
	{"", 0},
}

func TestLongestValidParentheses(t *testing.T) {
	for _, testCase := range testCases {
		result := longestValidParentheses(testCase.s)
		if result != testCase.want {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
