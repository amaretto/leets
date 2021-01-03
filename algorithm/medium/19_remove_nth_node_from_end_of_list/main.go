package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	var parray []*ListNode
	for p != nil {
		parray = append(parray, p)
		p = p.Next
	}

	// target is len(parray)-n
	befidx := len(parray) - n - 1
	aftidx := len(parray) - n + 1
	// in this case, list must be empty
	if befidx < -1 {
		return nil
	} else if befidx == -1 {
		return head.Next
	}
	// in this case, node in tail of list removed
	if aftidx > len(parray)-1 {
		parray[befidx].Next = nil
	} else {
		parray[befidx].Next = parray[aftidx]
	}

	return head
}
