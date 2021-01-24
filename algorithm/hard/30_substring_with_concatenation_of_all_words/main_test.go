package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	s     string
	words []string
	want  []int
}{
	{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int(nil)},
	{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
}

func TestFindSubString(t *testing.T) {
	for _, testCase := range testCases {

		result := findSubstring(testCase.s, testCase.words)

		if !reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
