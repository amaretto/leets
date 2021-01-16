package main

import "testing"

var testCases = []struct {
	head      *ListNode
	k int
	want     *ListNode
}{
	{&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4, &ListNode{5,nil}}}}}}, 2, &ListNode{2,&ListNode{1,&ListNode{4,&ListNode{3, &ListNode{5,nil}}}}}}},
}

func TestIsMatch(t *testing.T) {
	for _, testCase := range testCases {
		result := isMatch(testCase.Input, testCase.Expression)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}

}
