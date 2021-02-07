package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	head *ListNode
	k    int
	want *ListNode
}{
	{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}},
	//{&ListNode{1, nil}, 1, &ListNode{1, nil}},
}

func TestReverseKGroup(t *testing.T) {
	for _, testCase := range testCases {
		result := reverseKGroup(testCase.head, testCase.k)
		fmt.Println("head")
		printList(testCase.head)
		fmt.Println("result")
		printList(result)
		if reflect.DeepEqual(result, testCase.want) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}

}

func printList(ln *ListNode) {
	for ln != nil {
		fmt.Println(ln, ln.Val, ln.Next)
		ln = ln.Next
	}
}

func isSame(r, w *ListNode) {

}
