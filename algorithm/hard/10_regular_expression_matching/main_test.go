package main

import "testing"

var testCases = []struct {
	Input      string
	Expression string
	Output     bool
}{
	//	{"aa", "a", false},
	//	{"a", "a*", true},
	//	{"ab", ".*", true},
	//	{"abc", ".*c", true},
	//	{"aab", "c*a*b", true},
	//	{"aaa", "a*a", true},
	{"aaa", "ab*a*c*a", true},
	//	{"mississippi", "mis*is*p*", false},
	//	{"mississippi", "mis*is*ip*i", true},
	//	{"aa", "a*", true},
	//	{"ab", ".*c", false},
}

func TestIsMatch(t *testing.T) {
	for _, testCase := range testCases {
		result := isMatch(testCase.Input, testCase.Expression)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}

}
