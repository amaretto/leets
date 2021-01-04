package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []string
}{
	{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	{2, []string{"(())", "()()"}},
	{1, []string{"()"}},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, testCase := range testCases {
		result := generateParenthesis(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
