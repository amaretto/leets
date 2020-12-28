package main

import "testing"

var testCases = []struct {
	Nums   []int
	Output []int
}{
	{[]int{1, 1, 2}, []int{1, 2}},
	{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
}

func TestDeleteDuplicates(t *testing.T) {
	var input, output, pre, tmp *ListNode

	// build input
	for _, testCases := range testCase.Input {
		for _, num := range testCase.Nums {
			tmp := ListNode{num, nil}
			if pre != nil {
				pre.Next = tmp
			} else {
				input = tmp
			}
			pre = tmp
		}
	}
	// process
	deleteDuplicates(input)

	// build output
	tmp, pre = nil, nil
	for _, testCases := range testCase.Output {
		for _, num := range testCase.Nums {
			tmp := ListNode{num, nil}
			if pre != nil {
				pre.Next = tmp
			} else {
				output = tmp
			}
			pre = tmp
		}
	}

	for i, j := input, output; i.Next == nil && j.Next == nil; i, j = i.Next, j.Next {
		if i.Val != j.Val {
			t.Errorf("invalid result. expect:%v, actual:%v", j.Val, i.Val)
		}
	}
}
