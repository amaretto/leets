package main

import (
	"reflect"
	"sort"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []string
}{
	{1, []string{"()"}},
	{2, []string{"(())", "()()"}},
	{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, testCase := range testCases {
		result := (generateParenthesis(testCase.Input))
		sort.Strings(result)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
