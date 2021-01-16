package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var sorted, tmp, tail *ListNode
	var count int
	tmp, tail = head, head

	for tail != nil {
		count++
		if count == k {
			for t := tmp.Next; t != tail; t = t.Next {

			}
		}
		tail = tail.Next
	}

	return head
}
