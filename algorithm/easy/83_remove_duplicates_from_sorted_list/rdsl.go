package main

import "fmt"

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// test case a
	//c := ListNode{2, nil}
	//b := ListNode{1, &c}
	//a := ListNode{1, &b}

	// test case b
	e := ListNode{3, nil}
	d := ListNode{3, &e}
	c := ListNode{2, &d}
	b := ListNode{1, &c}
	a := ListNode{1, &b}

	// test case c

	deleteDuplicates(nil)

	for i := &a; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	pre := head
	for i := pre.Next; i != nil; i = i.Next {
		if pre.Val == i.Val {
			pre.Next = i.Next
		} else {
			pre = i
		}
	}
	return head
}
