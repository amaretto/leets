package main

import (
	"testing"
)

var testCases = []struct {
	input []*ListNode
	want  *ListNode
}{
	{
		[]*ListNode{&ListNode{1, &ListNode{4, &ListNode{5, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{2, &ListNode{6, nil}}},
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
	},
	{
		[]*ListNode{},
		nil,
	},
	{
		[]*ListNode{nil},
		nil,
	},
}

func TestIsMatch(t *testing.T) {
	for _, testCase := range testCases {
		result := mergeKLists(testCase.input)
		for i, j := result, testCase.want; i != nil && j != nil; i, j = i.Next, j.Next {
			if (i == nil || j == nil) || i.Val != j.Val {
				t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
			}
		}
	}
}
