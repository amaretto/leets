package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head, head.Next, head.Next.Next = head.Next, head.Next.Next, head
	swapPairs(head.Next.Next)
	return head
}
