package main

import "testing"

var testCases = []struct {
	Input      string
	Expression string
	Output     bool
}{
	{"aa", "a", false},
	{"a", "a*", true},
	{"ab", ".*", true},
	{"abc", ".*c", true},
	{"aab", "c*a*b", true},
	{"mississippi", "mis*is*p*", false},
	{"aa", "a*", true},
	{"ab", ".*c", false},
}

func TestIsMatch(t *testing.T) {
	for _, testCase := range testCases {
		result := isMatch(testCase.Input, testCase.Expression)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}

}
