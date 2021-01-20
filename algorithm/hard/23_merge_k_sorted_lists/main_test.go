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

func TestMergeKLists(t *testing.T) {
	for _, testCase := range testCases {
		result := mergeKLists(testCase.input)
		for i, j := result, testCase.want; i != nil && j != nil; i, j = i.Next, j.Next {
			if (i == nil || j == nil) || i.Val != j.Val {
				t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
			}
		}
	}
}

func BenchmarkMergeKLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lnl := []*ListNode{&ListNode{1, &ListNode{4, &ListNode{5, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{2, &ListNode{6, nil}}}
		mergeKLists(lnl)
	}
}

var testCasesImprove = []struct {
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

func TestMergeKListsImprove(t *testing.T) {
	for _, testCase := range testCasesImprove {
		result := mergeKListsImprove(testCase.input)
		for i, j := result, testCase.want; i != nil && j != nil; i, j = i.Next, j.Next {
			if (i == nil || j == nil) || i.Val != j.Val {
				t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
			}
		}
	}
}

func BenchmarkMergeKListsImprove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lnl := []*ListNode{&ListNode{1, &ListNode{4, &ListNode{5, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{2, &ListNode{6, nil}}}
		mergeKListsImprove(lnl)
	}
}
