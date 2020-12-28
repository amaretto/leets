package main

import "fmt"

// ListNode represents the ListNode of the numbers
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// init
	// case 1
	//	l1 := &ListNode{2, nil}
	//	l2 := &ListNode{5, nil}
	//
	//	l1.Next = &ListNode{4, nil}
	//	l1.Next.Next = &ListNode{3, nil}
	//
	//	l2.Next = &ListNode{6, nil}
	//	l2.Next.Next = &ListNode{4, nil}
	// case 2
	//	l1 := &ListNode{5, nil}
	//	l2 := &ListNode{5, nil}
	// case 3
	//	l1 := &ListNode{1, nil}
	//	l1.Next = &ListNode{8, nil}
	//	l2 := &ListNode{0, nil}
	// case 4
	l1 := &ListNode{1, nil}
	l2 := &ListNode{9, nil}
	l2.Next = &ListNode{9, nil}

	l := addTwoNumbers(l1, l2)
	printListNode(l)
}

func printListNode(l *ListNode) {
	tmp := l
	for {
		fmt.Println(tmp.Val)
		if tmp.Next == nil {
			break
		} else {
			tmp = tmp.Next
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var adv, l1end, l2end bool
	var sum int

	var l1tmp, l2tmp, tmp, result *ListNode

	l1tmp = l1
	l2tmp = l2

	for !l1end || !l2end {
		if adv {
			sum++
			adv = false
		}

		if !l1end && !l2end {
			sum += l1tmp.Val + l2tmp.Val
		} else if !l1end {
			sum += l1tmp.Val
		} else {
			sum += l2tmp.Val
		}

		if sum >= 10 {
			adv = true
			sum -= 10
		}

		if result == nil {
			result = &ListNode{sum, nil}
			tmp = result
		} else {
			tmp.Next = &ListNode{sum, nil}
			tmp = tmp.Next
		}

		if !l1end {
			if l1tmp.Next != nil {
				l1tmp = l1tmp.Next
			} else {
				l1end = true
			}
		}
		if !l2end {
			if l2tmp.Next != nil {
				l2tmp = l2tmp.Next
			} else {
				l2end = true
			}
		}
		sum = 0
	}

	if adv {
		tmp.Next = &ListNode{1, nil}
	}

	return result
}
