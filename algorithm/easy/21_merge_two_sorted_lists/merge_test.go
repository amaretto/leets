package main

import (
	"testing"
)

var testCases = []struct {
	L1     []int
	L2     []int
	Output []int
}{
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
}

func TestMerge(t *testing.T) {
	var l1, l2, tmp, newNode *ListNode
	for _, testCase := range testCases {
		// make L1
		for i, n := range testCase.L1 {
			if i == 0 {
				l1 = &ListNode{n, nil}
				tmp = l1
			} else {
				newNode = &ListNode{n, nil}
				tmp.Next = newNode
				tmp = newNode
			}
		}

		// make L2
		for i, n := range testCase.L2 {
			if i == 0 {
				l2 = &ListNode{n, nil}
				tmp = l2
			} else {
				newNode = &ListNode{n, nil}
				tmp.Next = newNode
				tmp = newNode
			}
		}

		result := mergeTwoLists(l1, l2)
		for _, n := range testCase.Output {
			//fmt.Println(n, ",", result.Val)
			if n == result.Val {
				result = result.Next
			} else {
				t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
			}
		}
	}
}
