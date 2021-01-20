package main

import (
	"container/heap"
	"fmt"
)

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

/* ############################### */
// this code is reffering to discussion in leetcodes
type ListNodeList []*ListNode

func (lnl ListNodeList) Len() int           { return len(lnl) }
func (lnl ListNodeList) Less(i, j int) bool { return lnl[i].Val < lnl[j].Val }
func (lnl ListNodeList) Swap(i, j int)      { lnl[i], lnl[j] = lnl[j], lnl[i] }

func (lnl *ListNodeList) Push(x interface{}) {
	*lnl = append(*lnl, x.(*ListNode))
}

func (lnl *ListNodeList) Pop() interface{} {
	old := *lnl
	n := len(old)
	x := old[n-1]
	*lnl = old[0 : n-1]
	return x
}

func mergeKListsImprove(lists []*ListNode) *ListNode {
	h := &ListNodeList{}

	for _, ln := range lists {
		if ln != nil {
			heap.Push(h, ln)
		}
	}

	if h.Len() == 0 {
		return nil
	}

	head := heap.Pop(h).(*ListNode)
	tail := head

	if head.Next != nil {
		heap.Push(h, head.Next)
	}

	for h.Len() > 0 {
		current := heap.Pop(h).(*ListNode)

		tail.Next = current
		tail = current

		if current.Next != nil {
			heap.Push(h, current.Next)
		}

	}

	return head
}
