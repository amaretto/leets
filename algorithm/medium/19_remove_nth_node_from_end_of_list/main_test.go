package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Target int
	Output []int
}{
	{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
	{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	{[]int{1}, 1, []int(nil)},
	{[]int{1, 2}, 1, []int{1}},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, testCase := range testCases {
		// create linked list
		var head, tmp *ListNode
		for i, num := range testCase.Input {
			if i == 0 {
				tmp = &ListNode{Val: num, Next: nil}
				head = tmp
			} else {
				tmp.Next = &ListNode{Val: num, Next: nil}
				tmp = tmp.Next
			}
		}

		result := removeNthFromEnd(head, testCase.Target)

		// create result array
		var rarray []int
		for result != nil {
			rarray = append(rarray, result.Val)
			result = result.Next
		}

		if !reflect.DeepEqual(rarray, testCase.Output) {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, rarray)
		}
	}
}
