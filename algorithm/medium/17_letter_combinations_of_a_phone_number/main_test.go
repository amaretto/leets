package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output []string
}{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"", []string(nil)},
	{"2", []string{"a", "b", "c"}},
	{"8", []string{"t", "u", "v"}},
}

func TestThreeSumClosest(t *testing.T) {
	for _, testCase := range testCases {
		result := letterCombinations(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
