package main

import "fmt"

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tmp, tmp1, tmp2 *ListNode

	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// init
	if l1.Val <= l2.Val {
		result = l1
		tmp = l1
		tmp1 = l1.Next
		tmp2 = l2
	} else {
		result = l2
		tmp = l2
		tmp1 = l1
		tmp2 = l2.Next
	}

	for {
		if tmp1 == nil {
			tmp.Next = tmp2
			break
		} else if tmp2 == nil {
			tmp.Next = tmp1
			break
		}
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp = tmp2
			tmp2 = tmp2.Next
		}
	}

	for tmp := result; ; tmp = tmp.Next {
		fmt.Println(tmp.Val)
		if tmp.Next == nil {
			break
		}
	}

	return result
}
