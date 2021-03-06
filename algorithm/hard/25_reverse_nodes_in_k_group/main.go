package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var count int
	tmp := head
	var ahead, ghead, ptail, pre, next *ListNode

	if head.Next == nil || k == 1 {
		return head
	}

	for tmp != nil {
		next = tmp.Next
		if count%k == 0 { // head of group
			ghead = tmp
			if ptail != nil {
				ptail.Next = tmp
			}
			if !hasGroup(tmp, k) {
				return ahead
			}
		} else if count%k == k-1 { // tail of group
			ghead.Next = tmp.Next
			if ptail != nil {
				ptail.Next = tmp
			}
			tmp.Next = pre
			if ahead == nil {
				ahead = tmp
			}
			ptail = ghead
		} else {
			tmp.Next = pre
		}
		pre = tmp
		tmp = next
		count++
	}

	return ahead
}

func hasGroup(head *ListNode, k int) bool {
	tmp := head
	for i := 0; i < k; i++ {
		if tmp == nil {
			return false
		}
		tmp = tmp.Next
	}
	return true
}
