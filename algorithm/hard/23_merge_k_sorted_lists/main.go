package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{1, nil}
	b := ListNode{4, nil}
	a.Next = &b
	c := ListNode{5, nil}
	b.Next = &c

	d := ListNode{1, nil}
	e := ListNode{3, nil}
	d.Next = &e
	f := ListNode{4, nil}
	e.Next = &f

	g := ListNode{2, nil}
	h := ListNode{6, nil}
	g.Next = &h
	printll(mergeKLists([]*ListNode{&a, &d, &g}))

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
