package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	b := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	c := &ListNode{2, &ListNode{6, nil}}

	printll(mergeKLists([]*ListNode{a, b, c}))

}

func printll(l *ListNode) {
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, tmp, minNode *ListNode
	var minIdx int
	for !isEmpty(lists) {
		min := 100000
		for i, n := range lists {
			if n == nil {
				continue
			}
			if n.Val < min {
				min = n.Val
				minNode = n
				minIdx = i
			}
		}
		if head == nil {
			head = minNode
		} else {
			tmp.Next = minNode
		}
		tmp = minNode
		lists[minIdx] = minNode.Next
	}

	return head
}

func isEmpty(lists []*ListNode) bool {
	for _, n := range lists {
		if n != nil {
			return false
		}
	}
	return true
}
